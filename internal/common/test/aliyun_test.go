package test

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/common/oss/aliyun"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	endPoin         = ``
	accessKeyID     = ``
	accessKeySecret = ``
	bucketName      = ``
)

func TestAliOssVersion(t *testing.T) {
	fmt.Println("OSS Go SDK Version: ", oss.Version)
}

func TestListBucketNames(t *testing.T) {
	oss, err := aliyun.NewAliyunOss(endPoin, accessKeyID, accessKeySecret)
	if err != nil {
		fmt.Println(err)
		return
	}

	assert.NotNil(t, oss)
	names, err := oss.ListBucketNames()
	assert.Nil(t, err)
	for _, e := range names {
		fmt.Println(e)
	}
}

func TestCreateBucket(t *testing.T) {
	oss, err := aliyun.NewAliyunOss(endPoin, accessKeyID, accessKeySecret)
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.NotNil(t, oss)
	oss.CreateBucket(uuid.New().String())
}

func TestDeleteBucket(t *testing.T) {
	oss, err := aliyun.NewAliyunOss(endPoin, accessKeyID, accessKeySecret)
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.NotNil(t, oss)
	oss.DeleteBucket("8f1ffd12-f5d8-4bef-8c68-9389ae58e15a")
}

func TestUploadFileToOss(t *testing.T) {
	oss, err := aliyun.NewAliyunOss(endPoin, accessKeyID, accessKeySecret)
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.NotNil(t, oss)
	testStr := "test string"
	err = oss.UploadBytes("testString", bucketName, []byte(testStr))
	assert.Nil(t, err)
}
