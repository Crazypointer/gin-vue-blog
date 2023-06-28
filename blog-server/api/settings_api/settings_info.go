package settings_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/res"
)

// 配置信息

type SettingsUri struct {
	Name string `uri:"name"`
}

// 实现 SettingsApi 的 SettingsInfoView 接口 查看网站信息

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	default:
		res.FailWithMessage("没有对应的配置信息", c)
	}
}
