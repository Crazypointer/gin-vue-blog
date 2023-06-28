package user_api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
	"server/plugins/email"
	"server/utils"
	"server/utils/jwts"
	"server/utils/pwd"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Password string  `json:"password"` // 传入的新密码        binding:"required" msg:"密码不能为空"
	Code     *string `json:"code"`     //binding:"required" msg:"验证码不能为空"
}

func (UserApi) UserBindEmailView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	//	用户绑定邮箱
	//后台会给这个邮箱发送一个验证码
	var cr BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	session := sessions.Default(c)
	if cr.Code == nil {
		//第一次 发送验证码
		//生成4位验证码
		code := utils.Code(4)
		//将验证码存入session
		session.Set("valid_code", code)
		err = session.Save()
		if err != nil {
			res.FailWithMessage("存入session失败", c)
			return
		}
		//发送
		err = email.NewCode().Send(cr.Email, "验证码是："+code)
		if err != nil {
			res.FailWithMessage("发送验证码失败", c)
			return
		}
		res.OkWithMessage("验证码已发送，请查收", c)
	} else {
		code := session.Get("valid_code")
		//第二次 用户输入验证码、密码
		//验证验证码
		if *cr.Code != code {
			res.FailWithMessage("验证码错误", c)
			return
		}
		var user models.UserModel
		err = global.DB.Take(&user, claims.UserID).Error
		if err != nil {
			res.FailWithMessage("用户不存在", c)
			return
		}
		if len(cr.Password) < 4 {
			res.FailWithMessage("密码强度过低", c)
			return
		}
		hashPwd := pwd.HashPwd(cr.Password)
		//修改用户邮箱
		err := global.DB.Model(&user).Updates(map[string]any{
			"email":    cr.Email,
			"password": hashPwd,
		}).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("邮箱绑定失败", c)
			return
		}
		//完成绑定
		res.OkWithMessage("邮箱绑定成功", c)
	}

}
