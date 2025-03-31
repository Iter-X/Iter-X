package bo

type InitUploadRequest struct {
	Filename string
}

type InitUploadReply struct {
	UploadID   string `json:"uploadId"`
	BucketName string `json:"bucketName"`
	ObjectKey  string `json:"objectKey"`
}

type GenerateUploadPartURLRequest struct {
	UploadID   string `json:"uploadId"`
	PartNumber int    `json:"partNumber"`
	ObjectKey  string `json:"objectKey"`
}

type GenerateUploadPartURLReply struct {
	UploadID       string `json:"uploadId"`
	BucketName     string `json:"bucketName"`
	ObjectKey      string `json:"objectKey"`
	PartNumber     int    `json:"partNumber"`
	UploadURL      string `json:"uploadUrl"`
	ExpirationTime int64  `json:"expirationTime"`
}

type UploadPart struct {
	PartNumber int    `json:"partNumber"` // Part number
	ETag       string `json:"eTag"`       // ETag value of the part's data
}

type CompleteUploadRequest struct {
	UploadID  string       `json:"uploadId"`
	ObjectKey string       `json:"objectKey"`
	Parts     []UploadPart `json:"parts"`
	FileSize  int64        `json:"fileSize"`
}

type CompleteUploadReply struct {
	Location   string `json:"location"`
	Bucket     string `json:"bucket"`
	Key        string `json:"key"`
	ETag       string `json:"eTag"`
	PrivateURL string `json:"privateURL"`
	PublicURL  string `json:"publicURL"`
	Expiration int64  `json:"expiration"`
	FileId     uint   `json:"fileId"`
}
