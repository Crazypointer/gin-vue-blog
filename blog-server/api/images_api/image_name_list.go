package images_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

type ImageResponse struct {
	ID   uint   `json:"id"`   // 图片id
	Path string `json:"path"` // 图片路径
	Name string `json:"name"` // 图片名称
}

// ImageNameListView 图片名称列表
// @Tags 图片管理
// @Summary 图片名称列表
// @Description 图片名称列表
// @Router /api/images_names [get]
// @Success 200 {object} res.Response{data=[]ImageResponse}
// @Produce json
func (ImagesApi) ImageNameListView(c *gin.Context) {
	var imageList []ImageResponse
	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&imageList)
	res.OkWithData(imageList, c)

}
