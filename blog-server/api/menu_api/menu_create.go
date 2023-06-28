package menu_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/ctype"
	"server/models/res"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"菜单标题不能为空" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"菜单路径不能为空" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"`                // 简介的切换时间
	BannerTime    int         `json:"banner_time" structs:"banner_time"`                    // 菜单图片的切换时间 为 0 表示不切换
	Sort          int         `json:"sort" binding:"required" msg:"请选择菜单序号" structs:"sort"` // 菜单的序号
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`                          // 具体的图片顺序
}

func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 重复值判断
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, "title = ?", cr.Title).RowsAffected
	if count > 0 {
		res.FailWithMessage("菜单标题已存在", c)
		return
	}
	//创建菜单数据
	menuModel := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	err = global.DB.Create(&menuModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加菜单失败", c)
		return
	}
	if len(cr.ImageSortList) == 0 {
		res.OkWithMessage("添加菜单成功", c)
		return
	}
	// 一个菜单可能存在多个轮播图，因此采用数组
	var menuBannerList []models.MenuBannerModel
	for _, sort := range cr.ImageSortList {
		//理论上 还需判断image_id对应的图片是否真实存在
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: sort.ImageID,
			Sort:     sort.Sort,
		})
	}
	//批量插入第三张表
	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单图片关联失败", c)
		return
	}
	res.OkWithMessage("菜单添加成功", c)
}
