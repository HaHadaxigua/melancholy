/******
** @date : 2/4/2021 12:40 AM
** @author : zrx
** @description:
******/
package msg

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/basic/model"
	"github.com/HaHadaxigua/melancholy/internal/file/envir"
	"mime/multipart"
	"path"
	"strings"
	"time"
)

type ReqFolderGetInfo struct {
	FolderID string `form:"folderID" json:"folderID" binding:"required"`
	UserID   int
}

type ReqFolderCreate struct {
	FolderName string `json:"filename"`
	ParentID   string `json:"parentID, omitempty"`

	ID     string `json:"id"`
	UserID int
}

type ReqFolderUpdate struct {
	FolderID string `json:"folderID" binding:"required"`
	NewName  string `json:"newName"`

	UserID int
}

type ReqFolderDelete struct {
	FolderID string `json:"folderID" binding:"required"`

	UserID int
}

type ReqFolderPatchDelete struct {
	FolderIDs []string `json:"folderIDs" binding:"required"`

	UserID int
}

// ReqFolderInclude 当前文件夹下包含的内容 todo: 整理请求体
type ReqFolderInclude struct {
	FolderID   string `json:"folderID" binding:"required"`
	SortedBy   string `json:"sortedBy"`
	Descending bool   `json:"descending"` // 是否是降序

	UserID int
}

type ReqFileCreate struct {
	ParentID   string `json:"parentID"`
	FileName   string `json:"fileName"`
	FileType   int    `json:"fileType"`   // 创建的文件类型
	Size       int    `json:"size"`       // 文件大小
	Address    string `json:"address"`    // 文件的oss地址
	Hash       string `json:"hash"`       // 文件hash
	BucketName string `json:"bucketName"` // 存储桶名
	Endpoint   string `json:"endpoint"`

	UserID int
}

//  Verify 验证请求是否合法
func (req ReqFileCreate) Verify() bool {
	if req.FileType == 0 {
		// 判断上传的文件类型id是否支持
		if ftid, ok := envir.MapFileTypeToID[path.Ext(req.FileName)]; !ok {
			return false
		} else {
			req.FileType = ftid
		}
	} else {
		if ftstr, ok := envir.MapFileTypeToStr[req.FileType]; !ok {
			return false
		} else {
			if !strings.HasSuffix(req.FileName, ftstr) {
				return false
			}
		}
	}
	return true
}

type ReqFileDelete struct {
	FileID string `json:"fileID" binding:"required"`
	UserID int
}

type ReqFilePatchDelete struct {
	FileIDs []string `json:"fileIDs" binding:"required"`

	UserID int
}

// ReqFileUpload 简单文件的上传请求
type ReqFileUpload struct {
	Data       []byte                `json:"data"`
	FileHeader *multipart.FileHeader `json:"fileHeader"`
	ParentID   string                `form:"parentID" json:"parentID"`
	FileType   int                   `form:"fileType" json:"fileType"`
	Encryption bool                  `form:"encryption" json:"encryption"`
	KeySecret  string                `form:"encryption" json:"keySecret"` // 进行加密的字符串

	UserID int
}

// ReqFileSearch 文件搜索
type ReqFileSearch struct {
	Fuzzy string     `json:"fuzzy"`                          // 通过name来搜索文件或者是文件夹
	Start *time.Time `json:"start" time_format:"2006-01-02"` // 文件的最早更新时间
	End   *time.Time `json:"end" time_format:"2006-01-02"`   // 文件的最后更新时间

	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	UserID int
}

// ReqFolderListFilter 列出文件夹
type ReqFolderListFilter struct {
	FolderID string `json:"folderID"`

	UserID int
}

// ReqFileListFilter 列出指定文件夹中的文件
type ReqFileListFilter struct {
	FolderID string `form:"folderID" json:"folderID"`

	UserID int
}

// ReqFileDownload 下载文件的请求
type ReqFileDownload struct {
	FileID string `form:"fileID" json:"fileID"`

	UserID int
}

// ReqDeleteInIntegration 文件夹和文件的整合删除方法
type ReqDeleteInIntegration struct {
	FolderIDs []string `json:"folderIDs"`
	FileIDs   []string `json:"fileIDs"`

	UserID int
}

//=====================分片上传方法====================================

// ReqFileMultiCheck 检查文件上传情况
type ReqFileMultiCheck struct {
	Hash     string `form:"hash" json:"hash"` // 要上传的文件hash
	Filename string `form:"filename" json:"filename"`

	UserID int
}

// ReqFileMultiUpload 上传文件分片
type ReqFileMultiUpload struct {
	Filename  string `form:"filename" json:"filename"`   // 文件名称
	Hash      string `form:"hash" json:"hash"`           // 文件hash， 根据文件hash找到文件
	ChunkID   string `form:"chunkID" json:"chunkID"`     // 文件的分片id
	ChunkHash string `form:"chunkHash" json:"chunkHash"` // 分片的hash
	Total     int    `form:"total" json:"total"`         // 总共有多少个分片

	FileHeader *multipart.FileHeader `form:"file" json:"file"` // describes a file part of a multipart request.
	UserID     int

	MineType string `form:"mine_type" json:"mine_type"`
	Name     string `form:"name" json:"name"`
	Phase    string `form:"phase" json:"phase"`
	Size     int    `form:"size" json:"size"`
}

func (req ReqFileMultiUpload) String() string {
	return fmt.Sprintf("MineType: %v, Name:%v, Phase: %v, Size: %v", req.MineType, req.Name, req.Phase, req.Size)
}

// ReqFileMultiMerge 请求将分片文件合并
type ReqFileMultiMerge struct {
	Filename string `json:"filename"`
	Hash     string `json:"hash"`

	UserID int
}

// ReqFileMultiDownload 文件的分片下载
type ReqFileMultiDownload struct {
	Hash string `json:"hash"`

	UserID int
}

// ReqFindFileByType 找出图片
type ReqFindFileByType struct {
	FileType int `form:"fileType" json:"fileType" binding:"required"`
	Offset   int `form:"offset" json:"offset"`
	Limit    int `form:"limit" json:"limit"`

	UserID int
}

// ReqDocFile 创建文本文件的请求
type ReqDocFile struct {
	Name    string `json:"name"`         // 文件名，需要后缀
	Content string `json:"content"`      // 文本内容
	ID      string `form:"id" json:"id"` // 用于更新时的id

	UserID int
}

// ReqVideoFile 关于视频文件的请求
type ReqVideoFile struct {
	Name              string   `form:"name" json:"name"`                           // 带有后缀的文件名
	Title             string   `form:"title" json:"title"`                         // 视频标题
	Description       string   `form:"description" json:"description"`             // 视频描述
	CoverUrl          string   `form:"coverUrl" json:"coverUrl"`                   // 视频封面地址
	Area              string   `form:"area" json:"area"`                           // 地区
	Species           string   `form:"species" json:"species"`                     // 视频类型
	ProductionCompany string   `form:"productionCompany" json:"productionCompany"` // 制作公司
	Years             int      `form:"years" json:"years"`                         // 年份
	Duration          int      `form:"duration" json:"duration"`                   // 时长
	Tags              []string `form:"tags" json:"tags"`                           // 视频标签
	ID                string   `form:"id" json:"id"`                               // 用于更新时的id
	Size              int      `form:"size" json:"size"`                           // 文件大小
	Hash              string   `json:"hash"`                                       // 文件hash
	VideoID           string   `json:"videoID" form:"videoID"`                     // 视频文件ID
	Region            string   `json:"region"`                                     // 存储地区
	Bucket            string   `json:"bucket"`
	Endpoint          string   `json:"endpoint"`

	UserID int
}

// ReqMusicFile 关于音频文件的请求
type ReqMusicFile struct {
	Name     string   `json:"name" form:"name"`         // 歌名
	CoverUrl string   `json:"coverUrl" form:"coverUrl"` // 封面地址
	Duration int      `json:"duration" form:"duration"` // 时长
	Singer   string   `json:"singer" form:"singer"`     // 歌手
	Album    string   `json:"album" form:"album"`       // 专辑
	Years    int      `json:"years" form:"years"`       // 年份
	Species  string   `json:"species" form:"species"`   // 类型
	Tags     []string `json:"tags" form:"tags"`         // 音频标签
	ID       string   `json:"id" form:"id"`             // 对应的文件id
	Size     int      `json:"size" form:"size"`         // 文件大小
	Hash     string   `json:"hash"`                     // 文件hash
	MusicID  string   `json:"musicID" form:"musicID"`   // 音频文件id
	Region   string   `json:"region"`                   // 存储地区
	Bucket   string   `json:"bucket"`
	Endpoint string   `json:"endpoint"`

	UserID int
}

// ReqChunk 分片上传请求
type ReqChunk struct {
	Phase    string `json:"phase"`     // 所属阶段
	MimeType string `json:"mime_type"` // 文件类型
	Size     int    `json:"size"`      // 文件大小
	Name     string `json:"name"`      // 文件名
}

type RspChunkData struct {
	EndOffset int    `json:"end_offset"` // 每个切片的大小
	SessionID string `json:"session_id"` // 标识一个上传
}

// RspChunk 分片文件返回
type RspChunk struct {
	Status string       `json:"status"` // 成功需要返回success
	Data   RspChunkData `json:"data"`
}

// ReqGetUploadAddressAndToken 获取视频上传地址和token
type ReqGetUploadAddressAndToken struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Filename    string `json:"filename"`
	CoverURL    string `json:"coverURL"`
	Tags        string `json:"tags"`

	UserID int
	User   *model.User
}

// ReqRefreshAddressAndToken 刷新视频上传地址token
type ReqRefreshAddressAndToken struct {
	VideoID string `json:"videoID"`

	UserID int
	User   *model.User
}

// ReqGetMezzanineInfo 获取视频或音频文件的下载地址
type ReqGetMezzanineInfo struct {
	VideoID  string `json:"videoID"`  // 视频id

	UserID int
	User   *model.User
}

// ReqGetPlayInfo 获取视频播放地址
type ReqGetPlayInfo struct {
	VideoID  string `json:"videoID"`  // 视频id

	UserID int
	User   *model.User
}