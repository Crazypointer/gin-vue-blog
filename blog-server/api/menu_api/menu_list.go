package menu_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

func (MenuApi) MenuListView(c *gin.Context) {

	// MenuModel 菜单表
	// MenuBannerModel 菜单和背景图的连接表
	// BannerModel 背景图表

	// 先查菜单 菜单不需要分页
	var menuList []models.MenuModel
	// 需要菜单id
	var menuIDList []uint
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList)
	// 查连接表 查出菜单id所对应的背景图
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id in (?)", menuIDList)
	var menus []MenuResponse
	for _, menu := range menuList {
		// menu 是单个菜单
		//解决空数组返回null的问题
		//banners := []Banner{} 法1
		var banners []Banner = make([]Banner, 0) //法2
		// 在查询到的背景图列表中 找到当前菜单对应的所有背景图
		for _, menuBanner := range menuBanners {
			if menuBanner.MenuID == menu.ID {
				banners = append(banners, Banner{
					ID:   menuBanner.BannerID,
					Path: menuBanner.BannerModel.Path,
				})
			}
		}
		menus = append(menus, MenuResponse{
			MenuModel: menu,
			Banners:   banners,
		})
	}
	res.OkWithData(menus, c)
}
