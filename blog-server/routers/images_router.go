package routers

import "server/api"

func (router RouterGroup) ImagesRouter() {
	// 通过ApiGroupApp对象获得其 SettingsApi 接口的方法 SettingsInfoView
	app := api.ApiGroupApp.ImagesApi
	router.POST("images", app.ImageUploadView)        // 上传图片
	router.GET("images", app.ImageListView)           // 查询图片列表
	router.GET("images_names", app.ImageNameListView) // 查询图片列表
	router.DELETE("images", app.ImageRemoveView)      // 删除图片列表
	router.PUT("images", app.ImageUpdateView)         // 修改图片名称
}
