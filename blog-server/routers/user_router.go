package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"server/api"
	"server/middleware"
)

var store = cookie.NewStore([]byte("Hasda8da25eHGFTLOjh52"))

func (router RouterGroup) UserRouter() {
	// 设置session中间件
	router.Use(sessions.Sessions("sessionid", store))

	// 通过ApiGroupApp对象获得其 SettingsApi 接口的方法 SettingsInfoView
	app := api.ApiGroupApp.UserApi
	router.POST("email_login", app.EmailLoginView)                                //邮箱登录
	router.POST("users", middleware.JwtAdmin(), app.UserCreateView)               //添加用户
	router.POST("user_bind_email", middleware.JwtAuth(), app.UserBindEmailView)   //用户绑定邮箱
	router.POST("logout", middleware.JwtAuth(), app.UserLogOutView)               //修改用户信息
	router.PUT("user_role", middleware.JwtAdmin(), app.UserUpdateRoleView)        //修改用户权限
	router.PUT("user_password", middleware.JwtAuth(), app.UserUpdatePasswordView) //修改用户密码
	router.GET("users", middleware.JwtAuth(), app.UserListView)                   //查询用户列表
	router.DELETE("users", middleware.JwtAdmin(), app.UserRemoveView)             //添加用户
}
