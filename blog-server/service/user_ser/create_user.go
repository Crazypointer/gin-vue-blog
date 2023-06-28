package user_ser

import (
	"errors"
	"server/global"
	"server/models"
	"server/models/ctype"
	"server/utils/pwd"
)

// 头像
// 方案：1.默认头像 2.随机头像

var Avatar = "/upload/avatar/default.jpeg"

func (UserService) CreateUser(userName, nickname, password, email, ip string, role ctype.Role) error {
	// 判断用户名是否存在
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		return errors.New("用户名已存在")
	}

	// 对密码进行hash
	hashPassword := pwd.HashPwd(password)
	// 入库
	err = global.DB.Create(&models.UserModel{
		UserName:   userName,
		NickName:   nickname,
		Password:   hashPassword,
		Email:      email,
		Role:       role,
		Avatar:     Avatar,
		IP:         ip,
		Addr:       "内网地址",
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		return errors.New("创建用户失败")
	}
	return nil
}
