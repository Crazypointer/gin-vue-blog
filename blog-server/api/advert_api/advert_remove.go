package advert_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

// AdvertRemoveView 删除广告
// @Tags 广告管理
// @Summary 删除广告
// @Description 删除广告
// @Param data body models.RemoveRequest true "广告id列表"
// @Router /api/adverts [delete]
// @Success 200 {object} res.Response{data=string}
// @Produce json
func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var advertList []models.AdvertModel
	count := global.DB.Find(&advertList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("广告不存在", c)
		return
	}

	global.DB.Delete(&advertList)
	res.OkWithMessage(fmt.Sprintf("共删除：%d 条广告", count), c)
}
