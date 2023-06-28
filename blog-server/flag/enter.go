package flag

import (
	"flag"
	"github.com/fatih/structs"
)

type Option struct {
	DB   bool
	User string // 用户名 指令：-u admin 或者 -u user
}

// Parse 解析命令行参数
func Parse() Option {
	db := flag.Bool("db", false, "初始化数据库")
	user := flag.String("u", "", "创建用户")
	// 解析命令行参数写入注册的flag里
	flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
	}
}

// IsWebStop 判断是否停止web服务 只有当db和user都为空时才停止

func IsWebStop(option Option) (f bool) {
	maps := structs.Map(&option)
	for _, v := range maps {
		switch val := v.(type) {
		case string:
			if val != "" {
				f = true
			}
		case bool:
			if val {
				f = true
			}
		}

	}
	return f
}

// SwitchOption 根据命令执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
		return
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
		return
	}
	//sys_flag.Usage()
}
