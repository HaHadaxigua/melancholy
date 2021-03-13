package common

// 处理文件下载
type Downloader interface {
	// 下载文件
	Download() ([]byte, error)
	// 进行数据分片
 	SplitDate()([]byte, error)
}