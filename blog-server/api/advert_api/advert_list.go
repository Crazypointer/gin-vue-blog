package advert_api

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/models/res"
	"server/service/common"
	"strings"
)

// AdvertListView 广告列表
// @Tags 广告管理
// @Summary 广告列表
// @Description 广告列表
// @Param data query models.PageInfo false "查询参数"
// @Router /api/adverts [get]
// @Success 200 {object} res.Response{data=res.ListResponse[models.AdvertModel]}
// @Produce json
func (AdvertApi) AdvertListView(c *gin.Context) {
	var cr models.PageInfo //分页查询参数
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 规定 前端来的请求就返回查询到的全部信息 后台管理端来的信息 就只显示is_show=true这个字段
	// 判断referer中是否包含admin 包含则全部返回 不包含就只返回is_show=true
	referer := c.GetHeader("Referer")
	isShow := true
	if strings.Contains(referer, "admin") {
		// admin端发来的请求
		isShow = false
	}
	list, count, _ := common.ComList(models.AdvertModel{IsShow: isShow}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	res.OkWithList(list, count, c)
}
