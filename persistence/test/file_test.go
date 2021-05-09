package test

import (
	"archive/zip"
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)


// 创建空文件
func TestCreateEmptyFile(t *testing.T) {
	newFile, err := os.Create("test.txt")
	if err != nil {
		logrus.Warn(err)
		return
	}
	defer newFile.Close()
	logrus.Println(newFile)
}

// 裁剪文件， 会得到指定大小的文件
func TestTruncateFile(t *testing.T) {
	if err := os.Truncate("test.txt", 10); err != nil {
		logrus.Warn(err)
		return
	}
}

// 得到文件信息
func TestGetFileInfo(t *testing.T){
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		logrus.Warn(err)
		return
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.Mode()) // permissions
	fmt.Println(fileInfo.ModTime()) // last modified time
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.Sys())
}

// 重命名和移动
func TestRenameAndMove(t *testing.T) {
	originPath := "test.txt"
	newPath := "../test.txt"
	if err := os.Rename(originPath, newPath); err != nil {
		logrus.Warn(err)
		return
	}
}

// 删除文件
func TestRemoveFile(t *testing.T) {
	if err := os.Remove("test.txt"); err != nil {
		logrus.Warn(err)
		return
	}
}


// 打开和关闭文件
func TestOpenAndCloseFile(t *testing.T) {
	// 以只读方式打开文件
	file, err := os.Open("test.txt")
	if err != nil {
		logrus.Warn(err)
		return
	}
	file.Close()

	// 更多方式的打开文件 参数二：打开的属性，参数三：权限模式
	file, err = os.OpenFile("text.txt", os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	file.Close()

	fmt.Println(os.O_RDWR) // 读写
	fmt.Println(os.O_RDONLY) // 只读
	fmt.Println(os.O_WRONLY) // write only
	fmt.Println(os.O_CREATE) // create if not exist
}

// TestFileIsExist 文件是否存在
func TestFileIsExist(t *testing.T) {
	file, err := os.Stat("D:\\data\\pictures\\wallpaper\\pixiv\\82847627_p0.png")
	fmt.Println("大小",file.Size())
	if err != nil {
		if os.IsNotExist(err) {
			logrus.Fatal("file not exist")
		}
		logrus.Fatal(err)
	}
	logrus.Println(file.Name())
}

// TestSymbolicFile 测试软链接, 软链接在windows中无效
func TestSymbolicFile(t *testing.T) {
	// create link, 创建一个硬连接，一个文件会有两个文件名。改变一个文件的内容会影响另一个。删除和重命名则不会。
	err := os.Link("origin.txt", "original_also.txt")
	if err  != nil {
		log.Fatal(err)
	}
}

// TestCopyFile 测试复制文件
func TestCopyFile(t *testing.T) {
	originFile, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originFile.Close()

	newFile, err := os.Create("newFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesWritten, err := io.Copy(newFile, originFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(bytesWritten)
	// 刷新内存并写入磁盘
	if err = newFile.Sync(); err != nil {
		log.Fatal(err)
	}
}

// 写文件
func TestWriteFile(t *testing.T){
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes := []byte("hello")
	written, err := file.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(written)
}


// TestQuickWrite 快写文件
func TestQuickWrite(t *testing.T) {
	err := ioutil.WriteFile("newFile.txt", []byte("helo"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

// 缓存写
func TestCacheWrite(t *testing.T) {
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// 为文件创建缓存写
	bufferedWriter := bufio.NewWriter(file)
	// 写字节到buffer
	bytesWritten, err := bufferedWriter.Write([]byte{65,66,67})
	if err != nil {
		log.Fatal(err)
	}
	logrus.Println("bytes available ", bufferedWriter.Available())

	logrus.Println("bytes write", bytesWritten)

	// 缓存中的字节数
	unflushedBufferSize := bufferedWriter.Buffered()
	logrus.Println("bytes buffered ", unflushedBufferSize)

	// 写到硬盘
	bufferedWriter.Flush()

	bufferedWriter.Reset(bufferedWriter) // 丢弃没缓存的内容，（将缓存给另一个writer时很有用）
	bytesAliable := bufferedWriter.Available()
	logrus.Println("bytes available", bytesAliable)

	bufferedWriter = bufio.NewWriterSize(bufferedWriter, 8000) // 重新设置缓存的大小

	logrus.Println("bytes avaliable", bufferedWriter.Available())
}

// 读取最多N个字节
func TestReadMaxNBytes(t *testing.T) {
	file,  err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	slice := make([]byte, 20)
	// file.Read可以读取一个小文件到大的slice中
	bytesRead, err := file.Read(slice)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(bytesRead)
}

// TestReadNBytes 测试读取N个字节
func TestReadNBytes(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	byteSlice := make([]byte, 2)
	// 当文件字节数小于byteSlice时会返回错误
	numBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(numBytesRead)
}

// 读取全部字节
func TestReadAll(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(data))
}

// 将文件内容快读到内存
func TestQuickReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(data))
}

// 缓存读， 将内容读到缓存之中
func TestBufferRead(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	bufferReader := bufio.NewReader(file)

	//byteSlice := make([]byte,5)
	byteSlice, err := bufferReader.Peek(5)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range byteSlice {
		fmt.Println(e)
	}
	dataBytes, err := bufferReader.ReadBytes('\n') // 读取到分隔符
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(dataBytes))
}

// 压缩文件
func TestZipFile(t *testing.T){
	outFile, err := os.Create("out.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	zipWriter := zip.NewWriter(outFile)
	var filesToArchive = []struct{Name, Body string} {
		{"test1.txt", "String contents of file"},
		{"test2.txt", "\x61\x62\x63\n"},
	}

	for _, file := range filesToArchive {
		fileWriter, err := zipWriter.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = fileWriter.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := zipWriter.Close(); err != nil {
		log.Fatal(err)
	}
}

// 解压文件
func TestUnzip(t *testing.T) {

}