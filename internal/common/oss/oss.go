package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

type MelancholyOSS interface {
	GetClient() *oss.Client
	ListBucketNames() ([]string,error)
	CreateBucket(bucketName string) *oss.Bucket
	DeleteBucket(bucketName string) error
	UploadBytes(bucketName, objectName string, data []byte) error
	UploadFileStream(objectName, bucketName string, fd *os.File) error
}