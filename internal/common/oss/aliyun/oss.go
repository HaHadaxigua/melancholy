/******
** @date : 2/15/2021 8:15 PM
** @author : zrx
** @description:
******/
package aliyun

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

type AliyunOss struct {
	client *oss.Client `json:"client"`
}

func NewAliyunOss(endPoint, accessKey, accessSecret string) *AliyunOss {
	cli, err := oss.New(endPoint, accessKey, accessSecret,
		oss.Timeout(10, 120), oss.EnableCRC(true))
	if err != nil {
		return nil
	}

	return &AliyunOss{
		client: cli,
	}
}

func (ali AliyunOss) GetClient() *oss.Client {
	return ali.client
}


func (api AliyunOss) ListBucketNames()([]string,error){
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
		}else {
			break
		}
	}
	return ret, nil
}

// CreateBucket 创建存储桶
func (ali AliyunOss) CreateBucket(bucketName string) *oss.Bucket {
	if err := ali.client.CreateBucket(bucketName, oss.ACL(oss.ACLPublicReadWrite)); err != nil {
		return nil
	}
	bucket, err := ali.client.Bucket(bucketName)
	if err != nil {
		return nil
	}
	return bucket
}

// DeleteBucket 删除存储桶
func (ali AliyunOss) DeleteBucket(bucketName string) error {
	return ali.client.DeleteBucket(bucketName)
}

// UploadBytes 上传byte数组
func (ali AliyunOss) UploadBytes(bucketName, objectName string, data []byte) error {
	client := ali.GetClient()
	// 获取存储空间
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
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

// UploadFileStream 上传文件流
func (api AliyunOss) UploadFileStream(objectName, bucketName string, fd *os.File) error {
	client := api.GetClient()
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}
	defer fd.Close()
	err = bucket.PutObject(objectName, fd)
	return err
}
