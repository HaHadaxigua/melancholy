package persistence

import (
	"bufio"
	"fmt"
	"os"
)

// Persistence 文件持久化
type Persistence interface {
	// 保存文件：将文件持久化 p: 文件名，文件位置，要保存的数据
	SaveSimpleFile(filename, location string, data []byte) error
}

// ResourceManager: 资源管理模块 处理本地文件持久化相关工作
type ResourceManager struct {

}

// SaveFile 保存文件，需要给出文件名、文件位置、以及数据
func (r ResourceManager) SaveSimpleFile(filename, location string, data []byte) error {
	file, err := os.OpenFile(location+filename, os.O_WRONLY, 06666)
	if err != nil {
		return err
	}
	defer func() {
		file.Close()
	}()
	bufferWritter := bufio.NewWriter(file)
	bytesWritten, err := bufferWritter.Write(data)
	if err != nil {
		return err
	}
	if bytesWritten < 1 {
		return fmt.Errorf("write file from buffer to hard-disk-dirve")
	}
	//unflushedBufferedSize := bufferWritter.Buffered()
	if err := bufferWritter.Flush(); err != nil {
		return err
	}
	bufferWritter.Reset(bufferWritter) // 丢弃没缓存的内容
	return nil
}


