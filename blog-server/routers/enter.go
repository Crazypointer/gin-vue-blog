package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"server/global"
	"server/middleware"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	//静态文件
	router.Static("/uploads", "./uploads")
	//跨域
	router.Use(middleware.Cors())
	//swagger文档
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	//路由分组
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	//系统配置api
	routerGroupApp.SettingsRouter()
	//图片管理api
	routerGroupApp.ImagesRouter()
	//广告管理api
	routerGroupApp.AdvertRouter()
	//菜单管理api
	routerGroupApp.MenuRouter()
	//用户管理api
	routerGroupApp.UserRouter()
	//标签管理api
	routerGroupApp.TagRouter()
	//文章管理api
	routerGroupApp.ArticleRouter()
	return router
}
