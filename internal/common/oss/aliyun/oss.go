/******
** @date : 2/15/2021 8:15 PM
** @author : zrx
** @description:
******/
package aliyun

import (
	"bytes"
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/consts"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"os"
	"strings"
)

type AliyunOss struct {
	client *oss.Client `json:"client"`
}

func NewAliyunOss(endPoint, accessKey, accessSecret string) (*AliyunOss, error) {
	endPoint = fmt.Sprintf("%s%s", consts.Https, endPoint)
	cli, err := oss.New(endPoint, accessKey, accessSecret,
		oss.Timeout(10, 120), oss.EnableCRC(true))
	if err != nil {
		return nil, err
	}

	return &AliyunOss{
		client: cli,
	}, nil
}

func (ali AliyunOss) GetClient() *oss.Client {
	return ali.client
}

func (api AliyunOss) ListBucketNames() ([]string, error) {
	var ret []string
	marker := ""
	for {
		lsRes, err := api.client.ListBuckets(oss.Marker(marker))
		if err != nil {
			return nil, err
		}
		for _, bucket := range lsRes.Buckets {
			ret = append(ret, bucket.Name)
		}
		if lsRes.IsTruncated {
			marker = lsRes.NextMarker
		} else {
			break
		}
	}
	return ret, nil
}

// CreateBucket 创建存储桶
func (ali AliyunOss) CreateBucket(bucketName string) (*oss.Bucket, error) {
	if err := ali.client.CreateBucket(bucketName, oss.ACL(oss.ACLPublicReadWrite)); err != nil {
		return nil, err
	}
	bucket, err := ali.client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

// DeleteBucket 删除存储桶
func (ali AliyunOss) DeleteBucket(bucketName string) error {
	return ali.client.DeleteBucket(bucketName)
}

// UploadBytes 上传byte数组
func (ali AliyunOss) UploadBytes(bucketName, objectName string, data []byte) error {
	client := ali.GetClient()

	// 判断存储桶是否存在
	isExist, err := client.IsBucketExist(bucketName)
	if err != nil {
		return err
	}

	var bucket *oss.Bucket

	// 不存在
	if !isExist {
		bucket, err = ali.CreateBucket(bucketName)
		if err != nil {
			return err
		}
	} else {
		bucket, err = client.Bucket(bucketName)
		// 获取存储空间
		if err != nil {
			return err
		}
	}
	storageType := oss.ObjectStorageClass(oss.StorageStandard)
	objectAcl := oss.ObjectACL(oss.ACLPublicReadWrite)

	// 上传byte数组
	err = bucket.PutObject(objectName, bytes.NewReader(data), storageType, objectAcl)
	if err != nil {
		return err
	}
	return nil
}

// UploadString 上传字符串
func (ali AliyunOss) UploadString(bucketName, objectName string, content string) error {
	client := ali.GetClient()

	// 判断存储桶是否存在
	isExist, err := client.IsBucketExist(bucketName)
	if err != nil {
		return err
	}

	var bucket *oss.Bucket

	// 不存在
	if !isExist {
		bucket, err = ali.CreateBucket(bucketName)
		if err != nil {
			return err
		}
	} else {
		bucket, err = client.Bucket(bucketName)
		// 获取存储空间
		if err != nil {
			return err
		}
	}
	storageType := oss.ObjectStorageClass(oss.StorageStandard)
	objectAcl := oss.ObjectACL(oss.ACLPublicReadWrite)

	err = bucket.PutObject(objectName, strings.NewReader(content), storageType, objectAcl)
	if err != nil {
		return err
	}
	return nil
}

// UploadFileStream 上传文件流
func (ali AliyunOss) UploadFileStream(objectName, bucketName string, fd *os.File) error {
	client := ali.GetClient()
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}
	defer fd.Close()
	err = bucket.PutObject(objectName, fd)
	return err
}

// DownloadFileByStream 处理文件的下载，写到缓存中
func (ali AliyunOss) DownloadFileByStream(bucketName, objectName string) (*bytes.Buffer, error) {
	client := ali.GetClient()
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	body, err := bucket.GetObject(objectName)
	if err != nil {
		return nil, err
	}
	defer body.Close()
	// 下载文件到缓存
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, body)
	if err != nil {
		return buf, err
	}
	return nil, err
}


// DeleteObject 删除单个文件
func (ali AliyunOss) DeleteSingleObject(bucketName, objectName string) error {
	bucket, err := ali.client.Bucket(bucketName)
	if err != nil {
		return err
	}
	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	if err := bucket.DeleteObject(objectName); err != nil {
		return err
	}
	return nil
}
