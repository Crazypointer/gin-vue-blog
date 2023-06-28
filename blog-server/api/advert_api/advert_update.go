package advert_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

// AdvertUpdateView 更新广告
// @Tags 广告管理
// @Summary 更新广告
// @Description 更新广告
// @Param data body AdvertRequest true "广告更新信息"
// @Router /api/adverts/:id [put]
// @Success 200 {object} res.Response{data=string}
// @Produce json
func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	var cr AdvertRequest
	id := c.Param("id")
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 判断是否存在 唯一索引：title 即如果title存在则认为广告已存在
	// 使用钩子函数 在每一次入库之前 判断是否存在
	var advert models.AdvertModel
	err = global.DB.Take(&advert, id).Error
	if err != nil {
		res.FailWithMessage("广告不存在", c)
		return
	}

	// 结构体转map
	maps := structs.Map(&cr)
	//添加
	err = global.DB.Model(&advert).Updates(maps).Error
	if err != nil {
		global.Log.Error("修改广告失败", err)
		res.FailWithMessage("修改广告失败", c)
		return
	}
	res.OkWithMessage("修改广告成功", c)
}
