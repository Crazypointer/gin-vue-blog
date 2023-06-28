package routers

import (
	"server/api"
)

func (router RouterGroup) SettingsRouter() {
	// 通过ApiGroupApp对象获得其 SettingsApi 接口的方法 SettingsInfoView
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("settings/:name", settingsApi.SettingsInfoView)       //查询网站信息
	router.PUT("settings/:name", settingsApi.SettingsInfoUpdateView) //更新网站信息
}
