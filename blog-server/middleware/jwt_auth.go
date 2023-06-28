package middleware

import (
	"github.com/gin-gonic/gin"
	"server/models/ctype"
	"server/models/res"
	"server/service/redis_ser"
	"server/utils/jwts"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//如何判断是管理员还是普通用户
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		//查看redis中是否存在token
		if redis_ser.CheckLogout(token) {
			res.FailWithMessage("用户已注销,token失效", c)
			c.Abort()
			return
		}

		//将用户信息写入上下文
		c.Set("claims", claims)
	}
}

// JwtAdmin 管理员权限才能调用的中间件
func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//如何判断是管理员还是普通用户
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token解析失败", c)
			c.Abort()
			return
		}
		//查看redis中是否存在token
		if redis_ser.CheckLogout(token) {
			res.FailWithMessage("用户已注销,token失效", c)
			c.Abort()
			return
		}
		//登陆的用户
		if claims.Role != int(ctype.PermissionAdmin) {
			res.FailWithMessage("非管理员用户，权限错误", c)
			c.Abort()
			return
		}
		//将用户信息写入上下文
		c.Set("claims", claims)
	}
}
