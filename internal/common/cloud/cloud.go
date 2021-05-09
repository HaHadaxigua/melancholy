package cloud

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
)

// 用于视频点播

var AliyunCloud MelancholyCloud

type MelancholyCloud interface {
	InitCloudClient(regionID string)
	InitVodClient() // 初始化视频点播客户端

	GetCloudClient() *sdk.Client                              // 获取云服务客户端
	GetVodClient() *vod.Client                                // 获取视频点播服务客户端
	UploadVideoByOriginOSSWithFilename(filename string) error // 通过oss原生sdk上传音视频文件
}
