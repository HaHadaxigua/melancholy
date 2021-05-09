# FileModule 文件模块

每一个用户都以一个root文件夹作为初始目录，与该用户有关的所有文件和文件夹都在这个初始文件夹下。

在创建新文件或者新文件夹的时候， 需要给出创建路径， 即当前是在哪个文件夹下创建的文件。

如果不给出路径， 则默认在初始文件夹下创建。

## 文件分块上传

前端需要将大文件分块，通知服务器文件是进行了分片的。

解决办法：

1. 前端发送check请求，询问服务器文件的上传情况
2. 判断是否存在名称为该文件hash的文件夹，如果存在返回其中的文件列表,如果存在一个文件名与文件夹名称一致的文件，说明已经上传成功；不存在，说明上传的文件为空;
3. 前端按照分片大小对文件进行分片，取得分片编号，和已经上传的文件列表进行对比，得到所有未上传的文件分片，将分片文件推送到服务器

## 21.02.14

文件模块的逻辑部分完成 todo:

1. 开通oss服务
2. 修改file数据结构 以支持oss
3. 上传文件到oss和从oss下载文件的方法
4. 上传文件接口， 下载文件接口

## 普通form表单获取数据的方法

```go
package demo

type Req struct {
	FileHeader *multipart.FileHeader `form:"file" json:"file"` // describes a file part of a multipart request.
	MineType   string                `form:"mine_type" json:"mine_type"`
	Name       string                `form:"name" json:"name"`
	Phase      string                `form:"phase" json:"phase"`
	Size       int                   `form:"size" json:"size"`
}

func uploadChunk(c *gin.Context) {
	var req msg.ReqFileMultiUpload
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErr(err))
		return
	}
	logrus.Info(req)
	return
}

```


