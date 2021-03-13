package test

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/common/oss/aliyun"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	endPoin         = `http://oss-cn-shanghai.aliyuncs.com`
	accessKeyID     = `LTAI4G9517RcKV7tRakHnHGh`
	accessKeySecret = `yxj7gbZ07J8lfZtvZw4BjdyldGUwpy`
	bucketName      = `test-upload-61227`
)

func TestListBucketNames(t *testing.T){
	oss := aliyun.NewAliyunOss(endPoin, accessKeyID, accessKeySecret)
	assert.NotNil(t, oss)
	names, err := oss.ListBucketNames()
	assert.Nil(t, err)
	for _, e := range names {
		fmt.Println(e)
	}
}

func TestCreateBucket(t *testing.T) {
	oss := aliyun.NewAliyunOss(endPoin, accessKeyID, accessKeySecret)
	assert.NotNil(t, oss)
	oss.CreateBucket(uuid.New().String())
}

func TestDeleteBucket(t *testing.T) {
	oss := aliyun.NewAliyunOss(endPoin, accessKeyID, accessKeySecret)
	assert.NotNil(t, oss)
	oss.DeleteBucket("8f1ffd12-f5d8-4bef-8c68-9389ae58e15a")
}

func TestUploadFileToOss(t *testing.T) {
	oss := aliyun.NewAliyunOss(endPoin, accessKeyID, accessKeySecret)
	assert.NotNil(t, oss)
	testStr := "test string"
	err := oss.UploadBytes("testString", bucketName,[]byte(testStr))
	assert.Nil(t, err)
}
