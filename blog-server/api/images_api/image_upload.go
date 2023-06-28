package images_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/res"
	"server/service"
	"server/service/image_ser"
)

// 上传图片 返回图片url

func (ImagesApi) ImageUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("文件不存在！", c)
		return
	}

	var resList []image_ser.FileUploadResponse

	for _, file := range fileList {
		// 上传文件
		serviceRes := service.ServiceApp.ImageService.ImageUploadService(file)
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		// 如果上传服务返回true 则完成了前置操作
		// 七牛云没有开启 则需要将图片上传到本地
		if !global.Config.QiNiu.Enable {
			// 上传图片到本服务器
			err = c.SaveUploadedFile(file, serviceRes.FileName)
			if err != nil {
				global.Log.Error(err)
				serviceRes.Msg = err.Error()
				serviceRes.IsSuccess = false
				resList = append(resList, serviceRes)
				continue
			}
		}

		// 七牛云开启 则存储七牛云存储后返回的地址
		resList = append(resList, serviceRes)
	}
	res.OkWithData(resList, c)
}
