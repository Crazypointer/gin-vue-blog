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

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required" msg:"旧密码不能为空"`
	NewPassword string `json:"new_password" binding:"required" msg:"新密码不能为空"`
}

// UserUpdatePasswordView 修改登陆人的密码
func (UserApi) UserUpdatePasswordView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	//绑定参数
	var cr UpdatePasswordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var user models.UserModel
	//查询数据库中的用户信息
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	fmt.Println(user.Password, cr.OldPassword)
	//判断旧密码是否正确
	if !pwd.ComparePwd(user.Password, cr.OldPassword) {
		res.FailWithMessage("旧密码错误", c)
		return
	}
	//更新密码
	hashPwd := pwd.HashPwd(cr.NewPassword)
	err = global.DB.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("密码修改失败", c)
		return
	}
	res.OkWithMessage("密码修改成功", c)

}
