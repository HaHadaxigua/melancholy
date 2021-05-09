package aliyun

import (
	"encoding/base64"
	"encoding/json"
	"github.com/HaHadaxigua/melancholy/internal/consts"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliyunCloud struct {
	AccessKeyID     string      `json:"accessKeyID"`
	AccessKeySecret string      `json:"accessKeySecret"`
	cloudClient     *sdk.Client `json:"client"`
	vodCloud        *vod.Client `json:"vodCloud"` // 视频点播客户端
}

func NewAliyunCloud(accessKeyID, accessKeySecret string) *AliyunCloud {
	return &AliyunCloud{
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
	}
}

func (ali AliyunCloud) InitCloudClient(regionID string) {
	client, err := sdk.NewClientWithAccessKey(regionID, ali.AccessKeyID, ali.AccessKeySecret)
	if err != nil {
		panic("initCloudClientFailed")
		return
	}
	ali.cloudClient = client
}

// NewVodClient 使用AK初始化VOD客户端
func (ali AliyunCloud) InitVodClient() {
	// 创建授权对象
	credential := &credentials.AccessKeyCredential{
		AccessKeyId:     ali.AccessKeyID,
		AccessKeySecret: ali.AccessKeySecret,
	}
	// 自定义config
	config := sdk.NewConfig()
	config.AutoRetry = true     // 失败是否自动重试
	config.MaxRetryTime = 3     // 最大重试次数
	config.Timeout = 3000000000 // 连接超时，单位：纳秒；默认为3秒
	// 创建vodClient实例
	vodClient, err := vod.NewClientWithOptions(consts.RegionID, config, credential)
	if err != nil {
		panic("initVodClientFailed")
		return
	}
	ali.vodCloud = vodClient
}

func (ali AliyunCloud) GetCloudClient() *sdk.Client {
	return ali.cloudClient
}

func (ali AliyunCloud) GetVodClient() *vod.Client {
	return ali.vodCloud
}

// CreateUploadVideo 获取视频上传地址和凭证
func (ali AliyunCloud) CreateUploadVideo() (*vod.CreateUploadVideoResponse, error) {
	req := vod.CreateCreateUploadVideoRequest()
	req.Title = "Sample Video Title"
	req.Description = "Sample Description"
	req.FileName = "/opt/video/sample/video_file.mp4"
	//request.CateId = "-1"
	req.CoverURL = "http://img.alicdn.com/tps/TB1qnJ1PVXXXXXCXXXXXXXXXXXX-700-700.png"
	req.Tags = "tag1,tag2"
	req.AcceptFormat = "JSON"
	return ali.GetVodClient().CreateUploadVideo(req)
}

type UploadAuthDTO struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
}
type UploadAddressDTO struct {
	Endpoint string
	Bucket   string
	FileName string
}

// InitOssClient 使用上传凭证和地址初始化OSS客户
func (ali AliyunCloud) InitOssClient(uploadAuthDTO UploadAuthDTO, uploadAddressDTO UploadAddressDTO) (*oss.Client, error) {
	client, err := oss.New(uploadAddressDTO.Endpoint,
		uploadAuthDTO.AccessKeyId,
		uploadAuthDTO.AccessKeySecret,
		oss.SecurityToken(uploadAuthDTO.SecurityToken),
		oss.Timeout(86400*7, 86400*7))
	return client, err
}

// 	uploadLocalFile 上传本地文件
func uploadLocalFile(client *oss.Client, uploadAddressDTO UploadAddressDTO, localFile string) error {
	// 获取存储空间。
	bucket, err := client.Bucket(uploadAddressDTO.Bucket)
	if err != nil {
		return err
	}
	// 上传本地文件。
	err = bucket.PutObjectFromFile(uploadAddressDTO.FileName, localFile)
	if err != nil {
		return err
	}
	return nil
}

// RefreshUploadVideo 刷新上传凭证
func(ali AliyunCloud) RefreshUploadVideo() (response *vod.RefreshUploadVideoResponse, err error) {
	request := vod.CreateRefreshUploadVideoRequest()
	request.VideoId = ""
	request.AcceptFormat = "JSON"
	return ali.vodCloud.RefreshUploadVideo(request)
}

// UploadVideoByOriginOSS 通过oss原生sdk上传视频，fileName: 需要上传到VOD的本地视频文件的完整路径
func (ali AliyunCloud) UploadVideoByOriginOSSWithFilename(filename string) error {
	// todo: 优化：需要创建一个新的文件上传服务器
	// 获取上传地址和凭证
	response, createUploadVideoErr := ali.CreateUploadVideo()
	if createUploadVideoErr != nil {
		return createUploadVideoErr
	}
	// 执行成功会返回VideoId、UploadAddress和UploadAuth
	//var videoId = response.VideoId
	var uploadAuthDTO UploadAuthDTO
	var uploadAddressDTO UploadAddressDTO
	var uploadAuthDecode, _ = base64.StdEncoding.DecodeString(response.UploadAuth)
	var uploadAddressDecode, _ = base64.StdEncoding.DecodeString(response.UploadAddress)
	json.Unmarshal(uploadAuthDecode, &uploadAuthDTO)
	json.Unmarshal(uploadAddressDecode, &uploadAddressDTO)
	// 使用UploadAuth和UploadAddress初始化OSS客户端
	ossClient, err := ali.InitOssClient(uploadAuthDTO, uploadAddressDTO)
	if err != nil {
		return err
	}
	// 上传文件，注意是同步上传会阻塞等待，耗时与文件大小和网络上行带宽有关
	//uploadLocalFile(ossClient, uploadAddressDTO, filename)
	go func() { // 开一个协程去处理
		uploadLocalFile(ossClient, uploadAddressDTO, filename)
	}()
	//MultipartUploadFile(ossClient, uploadAddressDTO, localFile)
	return nil
}
