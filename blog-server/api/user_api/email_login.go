package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
	"server/utils/jwts"
	"server/utils/pwd"
)

type EmailLogin struct {
	UserName string `json:"user_name" binding:"required" msg:"用户名不能为空"`
	Password string `json:"password" binding:"required" msg:"密码不能为空"`
}

func (UserApi) EmailLoginView(c *gin.Context) {
	var cr EmailLogin
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	fmt.Printf("%v", cr)
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ?", cr.UserName).Error
	if err != nil {
		global.Log.Warn("用户名不存在")
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 校验密码
	isCheck := pwd.ComparePwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("用户名或密码错误")
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NikeName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   int(userModel.ID),
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("生成token失败", c)
		return
	}
	res.OkWithData(token, c)
}
