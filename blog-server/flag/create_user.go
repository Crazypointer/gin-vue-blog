package flag

import (
	"fmt"
	"server/global"
	"server/models/ctype"
	"server/service/user_ser"
)

func CreateUser(permission string) {
	// 创建用户
	// 用户名 昵称 密码 确认密码 邮箱
	var (
		username   string
		nickname   string
		password   string
		rePassword string
		email      string
	)
	fmt.Print("请输入用户名：")
	fmt.Scan(&username)
	fmt.Print("请输入昵称：")
	fmt.Scanln(&nickname)
	fmt.Print("请输入密码：")
	fmt.Scan(&password)
	fmt.Print("请确认密码：")
	fmt.Scan(&rePassword)
	fmt.Print("请输入邮箱：")
	fmt.Scanln(&email)
	fmt.Println(username, nickname, password, rePassword, email)
	// 判断密码是否一致
	if password != rePassword {
		global.Log.Error("两次输入的密码不一致，请重新输入")
		return
	}
	// 判断权限
	role := ctype.PermissionUser
	if permission == "admin" {
		role = ctype.PermissionAdmin
	}
	//创建用户
	err := user_ser.UserService{}.CreateUser(username, nickname, password, email, "127.0.0.1", role)
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("用户：%s 创建成功", username)
}
