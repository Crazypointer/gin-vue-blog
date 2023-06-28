package tag_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

// AdvertModel 广告表

type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"请输入标题" structs:"title"` // 显示的标题
}

// TagCreateView 创建标签
// @Tags 标签管理
// @Summary 创建标签
// @Description 创建标签
// @Param data body TagRequest true "标签信息"
// @Router /api/tags [post]
// @Success 200 {object} res.Response{}
// @Failure 500 {object} res.Response{}
// @Produce json
func (TagApi) TagCreateView(c *gin.Context) {
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var tag models.TagModel
	err = global.DB.Take(&tag, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("该标签已存在", c)
		return
	}

	//添加
	err = global.DB.Create(&models.TagModel{
		Title: cr.Title,
	}).Error
	if err != nil {
		global.Log.Error("添加标签失败", err)
		res.FailWithMessage("添加标签失败", c)
		return
	}
	res.OkWithMessage("添加标签成功", c)
}
