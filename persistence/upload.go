package persistence

// 处理文件上传
type Uploader interface {
	Save(data []byte) error
}
