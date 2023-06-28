package routers

import "server/api"

func (router RouterGroup) ArticleRouter() {
	// 通过ApiGroupApp对象获得其 SettingsApi 接口的方法 SettingsInfoView
	app := api.ApiGroupApp.ArticleApi
	router.POST("articles", app.ArticleCreateView)
	router.GET("articles", app.ArticleListView)
	router.GET("articles/:id", app.ArticleSingleView)
	router.PUT("articles/:id", app.ArticleUpdateView)
	router.DELETE("articles", app.ArticleRemoveView)
}
