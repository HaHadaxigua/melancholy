/******
** @date : 2/15/2021 8:15 PM
** @author : zrx
** @description:
******/
package oss

import (
	"bytes"
	"github.com/HaHadaxigua/melancholy/internal/conf"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func GetOssClient() (*oss.Client, error) {
	client, err := oss.New(conf.C.Oss.EndPoint, conf.C.Oss.AccessKeyID, conf.C.Oss.AccessKeySecret,
		oss.Timeout(10, 120), oss.EnableCRC(true))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func UploadFile(objectName string, data []byte) {
	client, err := GetOssClient()
	if err != nil {
		return
	}
	bucket, err := client.Bucket(conf.C.Oss.BucketName)
	if err != nil {
		return
	}
	storageType := oss.ObjectStorageClass(oss.StorageStandard)
	objectAcl := oss.ObjectACL(oss.ACLPublicReadWrite)

	err = bucket.PutObject(objectName, bytes.NewReader(data), storageType, objectAcl)
	if err != nil {
		return
	}
}
