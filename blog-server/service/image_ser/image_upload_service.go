package image_ser

import (
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"path/filepath"
	"server/global"
	"server/models"
	"server/models/ctype"
	"server/plugins/qiniu"
	"server/utils"
	"strings"
)

type FileUploadResponse struct {
	ID        uint   `json:"id"` //返回上传图片的数据库id   99999999值为非法文件或文件大小超限制
	FileName  string `json:"fileName"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"` // 返回上传信息
}

// ImageUploadService 文件上传的方法

func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	res.FileName = file.Filename
	res.IsSuccess = false
	suffix := strings.ToLower(filepath.Ext(file.Filename)) // filepath.Ext(file.Filename) 获取文件后缀名 返回为：.png
	// 文件后缀 白名单判断
	if !utils.InList(suffix, global.ImageWhiteList) {
		res.ID = 99999999
		res.Msg = "非法文件"
		return
	}

	// 判断文件大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		res.ID = 99999999
		res.Msg = fmt.Sprintf("图片大小超限，上传失败！当前图片大小为%.2f，最大支持上传大小为：%d", size, global.Config.Upload.Size)
		return
	}

	//获取文件上传目标路径
	filePath := path.Join(global.Config.Upload.Path, file.Filename) //path.Join("uploads", file.Filename) 的作用是将字符串 "uploads" 和变量 file.Filename 连接起来形成一个路径字符串
	// 读取文件内容
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
	}
	byteData, err := io.ReadAll(fileObj)

	// 计算md5值
	imageHash := utils.Md5(byteData)
	//去数据库中查询md5值是否存在
	var bannerModel models.BannerModel
	err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
	if err == nil {
		// 找到了 返回数据库中的图片链接
		res.ID = bannerModel.ID
		res.FileName = bannerModel.Path
		res.Msg = "文件已存在"
		return
	}

	// 数据库中没有 需要上传到服务器端
	res.FileName = filePath
	fileType := ctype.Local
	res.IsSuccess = true
	res.Msg = "图片上传成功"

	// 七牛云开启 则上传到七牛云
	if global.Config.QiNiu.Enable {
		filePath, err = qiniu.UploadImage(byteData, file.Filename, global.Config.QiNiu.Prefix)
		if err != nil {
			global.Log.Error(err)
		}
		//添加响应信息
		res.FileName = filePath
		res.Msg = "文件已上传至七牛云"
		fileType = ctype.QiNiu
	}

	// 图片写入数据库
	global.DB.Create(&models.BannerModel{
		Path:      filePath,
		Hash:      imageHash,
		Name:      file.Filename,
		ImageType: fileType,
	})
	err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
	if err == nil {
		// 找到了 返回数据库中的图片链接
		res.ID = bannerModel.ID
	}
	return

}
