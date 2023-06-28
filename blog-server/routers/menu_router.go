package routers

import "server/api"

func (router RouterGroup) MenuRouter() {
	// 通过ApiGroupApp对象获得其 SettingsApi 接口的方法 SettingsInfoView
	app := api.ApiGroupApp.MenuApi
	router.POST("menus", app.MenuCreateView)        // 修改图片名称
	router.GET("menus", app.MenuListView)           // 查询菜单列表
	router.GET("menus_names", app.MenuNameListView) // 查询菜单名称列表
	router.PUT("menus/:id", app.MenuUpdateView)     // 更新菜单
	router.DELETE("menus", app.MenuRemoveView)      // 删除菜单
	router.GET("menus/:id", app.MenuDetailView)     // 获取菜单详情
}
