package menu_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	id := c.Param("id")
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}
	//先把之前的banner清空
	err = global.DB.Model(&menuModel).Association("Banners").Clear()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	//如果选择了banner，就把banner的数据添加进去
	if len(cr.ImageSortList) > 0 {
		//添加banner数据
		// MenuBannerModel 为中间表 用于连接菜单和背景图
		var bannerList []models.MenuBannerModel
		for _, v := range cr.ImageSortList {
			bannerList = append(bannerList, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: v.ImageID,
				Sort:     v.Sort,
			})
		}
		err = global.DB.Create(&bannerList).Error
		if err != nil {
			res.FailWithMessage("创建菜单图片失败", c)
			return
		}
	}

	//普通更新菜单数据
	// 结构体转map
	maps := structs.Map(&cr)
	//添加
	err = global.DB.Model(&menuModel).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改菜单失败", c)
		return
	}
	res.OkWithMessage("修改菜单成功", c)
}
