package advert_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

// AdvertModel 广告表

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`              // 显示的标题
	Href   string `json:"href" binding:"required,url" msg:"跳转图片链接非法" structs:"href"`         // 跳转链接
	Images string `json:"images" binding:"required,url" msg:"请输入一个合法的图片地址" structs:"images"` // 图片
	IsShow bool   `json:"is_show" msg:"请选择是否展示" structs:"is_show"`                           // 是否展示
}

// AdvertCreateView 创建广告
// @Tags 广告管理
// @Summary 创建广告
// @Description 创建广告
// @Param data body AdvertRequest true "广告信息"
// @Router /api/adverts [post]
// @Success 200 {object} res.Response{}
// @Failure 500 {object} res.Response{}
// @Produce json
func (AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 判断是否存在 唯一索引：title 即如果title存在则认为广告已存在
	// 使用钩子函数 在每一次入库之前 判断是否存在
	var advert models.AdvertModel
	err = global.DB.Take(&advert, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("广告已存在", c)
		return
	}

	//添加
	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error
	if err != nil {
		global.Log.Error("添加广告失败", err)
		res.FailWithMessage("添加广告失败", c)
		return
	}
	res.OkWithMessage("添加广告成功", c)
}
