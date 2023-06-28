package routers

import "server/api"

func (router RouterGroup) TagRouter() {
	// 通过ApiGroupApp对象获得其 SettingsApi 接口的方法 SettingsInfoView
	app := api.ApiGroupApp.TagApi
	router.POST("tags", app.TagCreateView)    // 创建标签
	router.GET("tags", app.TagListView)       // 获取标签列表
	router.PUT("tags/:id", app.TagUpdateView) // 更新标签
	router.DELETE("tags", app.TagRemoveView)  // 批量删除标签
}
