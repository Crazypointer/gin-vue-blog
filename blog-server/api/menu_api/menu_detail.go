package menu_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

func (MenuApi) MenuDetailView(c *gin.Context) {

	// MenuModel 菜单表
	// MenuBannerModel 菜单和背景图的连接表
	// BannerModel 背景图表

	// 先查菜单 需要菜单id 菜单不需要分页
	id := c.Param("id")
	var menuModel models.MenuModel
	err := global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}
	// 查连接表 查出菜单id所对应的背景图
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id = ?", id)

	var banners = make([]Banner, 0)
	for _, menuBanner := range menuBanners {
		if menuBanner.MenuID == menuModel.ID {
			banners = append(banners, Banner{
				ID:   menuBanner.BannerID,
				Path: menuBanner.BannerModel.Path,
			})
		}
	}
	menuResponse := MenuResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}

	res.OkWithData(menuResponse, c)
}
