package routers

import "server/api"

func (router RouterGroup) AdvertRouter() {
	// 通过ApiGroupApp对象获得其 SettingsApi 接口的方法 SettingsInfoView
	app := api.ApiGroupApp.AdvertApi
	router.POST("adverts", app.AdvertCreateView)
	router.GET("adverts", app.AdvertListView)
	router.PUT("adverts/:id", app.AdvertUpdateView)
	router.DELETE("adverts", app.AdvertRemoveView)
}
