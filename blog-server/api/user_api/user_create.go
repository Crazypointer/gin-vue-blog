package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/ctype"
	"server/models/res"
	"server/service/user_ser"
)

type UserCreateRequest struct {
	UserName string     `json:"user_name" binding:"required" msg:"用户名不能为空"` // 用户名
	NickName string     `json:"nick_name" binding:"required" msg:"昵称不能为空"`  // 昵称
	Password string     `json:"password" binding:"required" msg:"密码不能为空"`   // 密码`
	Role     ctype.Role `json:"role" binding:"required" msg:"权限不能为空"`       // 权限  1 管理员  2 普通用户  3 游客`
}

func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//创建用户
	err = user_ser.UserService{}.CreateUser(cr.UserName, cr.NickName, cr.Password, "", c.ClientIP(), cr.Role)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("创建用户失败", c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("用户：%s 创建成功", cr.UserName), c)
}
