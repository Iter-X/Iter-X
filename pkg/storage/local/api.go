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
)

func (l *Local) UploadHandler(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != l.uploadMethod {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	uploadID := r.URL.Query().Get("uploadID")
	partNumberStr := r.URL.Query().Get("partNumber")
	if uploadID == "" || partNumberStr == "" {
		http.Error(w, "missing uploadID or partNumber", http.StatusBadRequest)
		return
	}

	partNumber, err := strconv.Atoi(partNumberStr)
	if err != nil || partNumber <= 0 {
		http.Error(w, "invalid partNumber", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	l.rw.RLock()
	session, exists := l.uploads[uploadID]
	l.rw.RUnlock()

	if !exists {
		http.Error(w, "upload session not found", http.StatusNotFound)
		return
	}

	tempDir := filepath.Join(l.root, "tmp", uploadID)
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		http.Error(w, fmt.Sprintf("failed to create temp directory: %v", err), http.StatusInternalServerError)
		return
	}

	tempFile, err := os.CreateTemp(tempDir, fmt.Sprintf("part_%d", partNumber))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create temp file: %v", err), http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	hasher := md5.New()
	multiWriter := io.MultiWriter(tempFile, hasher)

	if _, err := io.Copy(multiWriter, r.Body); err != nil {
		http.Error(w, fmt.Sprintf("failed to write part data: %v", err), http.StatusInternalServerError)
		return
	}

	eTag := hex.EncodeToString(hasher.Sum(nil))

	l.rw.Lock()
	session.parts[partNumber] = tempFile.Name()
	l.rw.Unlock()

	w.Header().Set("ETag", eTag)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"uploadID":   uploadID,
		"partNumber": partNumber,
		"eTag":       eTag,
		"size":       hasher.Size(),
	})
}

func (l *Local) PreviewHandler(w http.ResponseWriter, r *http.Request) {
	objectKey := r.URL.Query().Get("objectKey")
	if objectKey == "" {
		http.Error(w, "missing objectKey", http.StatusBadRequest)
		return
	}

	filePath := objectKey
	if !strings.HasPrefix(filePath, l.root) {
		filePath = filepath.Join(l.root, filePath)
	}
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to open file: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	defer r.Body.Close()

	ext := filepath.Ext(file.Name())
	contentType := getContentType(ext)
	w.Header().Set("Content-Type", contentType)
	io.Copy(w, file)
	return
}

func getContentType(ext string) string {
	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".pdf":
		return "application/pdf"
	case ".mp4":
		return "video/mp4"
	case ".mp3":
		return "audio/mpeg"
	default:
		return "application/octet-stream"
	}
}
