package local

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/iter-x/iter-x/pkg/oss"
)

var _ oss.FileManager = (*local)(nil)

func NewLocalOSS(root, uploadMethod, uploadURL string) (oss.FileManager, error) {
	// 确保根目录存在
	if err := os.MkdirAll(root, 0755); err != nil {
		return nil, fmt.Errorf("failed to create root directory: %w", err)
	}
	return &local{
		root:         root,
		uploadURL:    uploadURL,
		uploadMethod: strings.ToUpper(uploadMethod),
		uploads:      make(map[string]*uploadSession),
	}, nil
}

type local struct {
	rw sync.RWMutex

	root, uploadMethod, uploadURL string
	uploads                       map[string]*uploadSession // uploadID -> session
}

type uploadSession struct {
	objectKey string
	parts     map[int]*os.File // partNumber -> temp file
	createdAt time.Time
}

func (l *local) uploadHandler(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != l.uploadMethod {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 获取必要参数
	uploadID := r.URL.Query().Get("uploadID")
	partNumberStr := r.URL.Query().Get("partNumber")
	if uploadID == "" || partNumberStr == "" {
		http.Error(w, "missing uploadID or partNumber", http.StatusBadRequest)
		return
	}

	// 转换partNumber为整数
	partNumber, err := strconv.Atoi(partNumberStr)
	if err != nil || partNumber <= 0 {
		http.Error(w, "invalid partNumber", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	// 获取上传会话（带读锁）
	l.rw.RLock()
	session, exists := l.uploads[uploadID]
	l.rw.RUnlock()

	if !exists {
		http.Error(w, "upload session not found", http.StatusNotFound)
		return
	}

	// 创建临时文件存储分片数据
	tempDir := filepath.Join(l.root, "tmp", uploadID)
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		http.Error(w, fmt.Sprintf("failed to create temp directory: %v", err), http.StatusInternalServerError)
		return
	}

	tempFile, err := os.CreateTemp(tempDir, fmt.Sprintf("part-%d-", partNumber))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create temp file: %v", err), http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	// 同时写入文件和计算ETag
	hasher := md5.New()
	multiWriter := io.MultiWriter(tempFile, hasher)

	if _, err := io.Copy(multiWriter, r.Body); err != nil {
		http.Error(w, fmt.Sprintf("failed to write part data: %v", err), http.StatusInternalServerError)
		return
	}

	// 获取ETag
	eTag := hex.EncodeToString(hasher.Sum(nil))

	// 将分片信息存入session（带写锁）
	l.rw.Lock()
	session.parts[partNumber] = tempFile
	l.rw.Unlock()

	// 设置响应头
	w.Header().Set("ETag", eTag)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// 返回JSON响应
	json.NewEncoder(w).Encode(map[string]interface{}{
		"uploadID":   uploadID,
		"partNumber": partNumber,
		"eTag":       eTag,
		"size":       hasher.Size(),
	})
}

func (l *local) generateObjectKey(originalName, group string) string {
	return fmt.Sprintf("%s/%s/%s/%d_%s", l.root, group, time.Now().Format("2006_01_02"), time.Now().UnixNano(), originalName)
}

func (l *local) InitiateMultipartUpload(originalName, group string) (*oss.InitiateMultipartUploadResult, error) {
	l.rw.Lock()
	defer l.rw.Unlock()

	objectKey := l.generateObjectKey(originalName, group)
	uploadID := generateUploadID(objectKey)

	l.uploads[uploadID] = &uploadSession{
		objectKey: objectKey,
		parts:     make(map[int]*os.File),
		createdAt: time.Now(),
	}

	return &oss.InitiateMultipartUploadResult{
		UploadID:   uploadID,
		BucketName: "local",
		ObjectKey:  objectKey,
	}, nil
}

func (l *local) GenerateUploadPartURL(uploadID, objectKey string, partNumber int, expires time.Duration) (*oss.UploadPartInfo, error) {
	// 本地实现不需要预签名URL，直接返回模拟信息
	return &oss.UploadPartInfo{
		UploadID:       uploadID,
		BucketName:     "local",
		ObjectKey:      objectKey,
		PartNumber:     partNumber,
		UploadURL:      fmt.Sprintf("local://%s/part/%d", objectKey, partNumber),
		ExpirationTime: time.Now().Add(expires).Unix(),
	}, nil
}

func (l *local) uploadPart(uploadID string, partNumber int, reader io.Reader) error {
	l.rw.Lock()
	defer l.rw.Unlock()

	session, exists := l.uploads[uploadID]
	if !exists {
		return fmt.Errorf("upload session not found")
	}

	// 创建临时目录存储分片
	tempDir := filepath.Join(l.root, "tmp", uploadID)
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}

	tempFile, err := os.CreateTemp(tempDir, fmt.Sprintf("part-%d-", partNumber))
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer tempFile.Close()

	if _, err := io.Copy(tempFile, reader); err != nil {
		return fmt.Errorf("failed to write part data: %w", err)
	}

	// 保存分片引用
	session.parts[partNumber] = tempFile

	return nil
}

func (l *local) CompleteMultipartUpload(uploadID, objectKey string, parts []oss.UploadPart) (*oss.CompleteMultipartUploadResult, error) {
	l.rw.Lock()
	defer l.rw.Unlock()

	defer func() {
		tempDir := filepath.Join(l.root, "tmp", uploadID)
		if err := os.RemoveAll(tempDir); err != nil {
			// TODO
		}
	}()

	session, exists := l.uploads[uploadID]
	if !exists {
		return nil, fmt.Errorf("upload session not found")
	}
	if objectKey != session.objectKey {
		return nil, fmt.Errorf("object key mismatch")
	}

	// 确保目标目录存在
	finalPath := filepath.Join(l.root, session.objectKey)
	if err := os.MkdirAll(filepath.Dir(finalPath), 0755); err != nil {
		return nil, fmt.Errorf("failed to create target directory: %w", err)
	}

	// 创建最终文件
	finalFile, err := os.Create(finalPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create final file: %w", err)
	}
	defer finalFile.Close()

	// 合并所有分片
	hasher := md5.New()
	multiWriter := io.MultiWriter(finalFile, hasher)

	for _, part := range parts {
		partFile, ok := session.parts[part.PartNumber]
		if !ok {
			return nil, fmt.Errorf("part %d not found", part.PartNumber)
		}

		if _, err := partFile.Seek(0, 0); err != nil {
			return nil, fmt.Errorf("failed to seek part file: %w", err)
		}

		if _, err := io.Copy(multiWriter, partFile); err != nil {
			return nil, fmt.Errorf("failed to merge part %d: %w", part.PartNumber, err)
		}
	}

	// 计算最终ETag
	eTag := hex.EncodeToString(hasher.Sum(nil))

	// 清理临时文件
	if err := os.RemoveAll(filepath.Join(l.root, "tmp", uploadID)); err != nil {
		return nil, fmt.Errorf("failed to clean temp files: %w", err)
	}

	// 从session中移除
	delete(l.uploads, uploadID)

	return &oss.CompleteMultipartUploadResult{
		Location:   finalPath,
		Bucket:     "local",
		Key:        session.objectKey,
		ETag:       eTag,
		PrivateURL: finalPath,
		PublicURL:  finalPath,
	}, nil
}

func (l *local) GeneratePublicURL(objectKey string, exp time.Duration) (string, error) {
	// 本地实现直接返回文件路径
	return filepath.Join(l.root, objectKey), nil
}

func generateUploadID(objectKey string) string {
	h := md5.New()
	h.Write([]byte(objectKey))
	h.Write([]byte(time.Now().String()))
	return hex.EncodeToString(h.Sum(nil))
}
