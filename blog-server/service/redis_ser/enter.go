package redis_ser

import (
	"server/global"
	"server/utils"
	"time"
)

var prefix = "logout_"

// 针对注销的操作

func Logout(token string, diff time.Duration) error {
	return global.Redis.Set(prefix+token, "", diff).Err()
}

// CheckLogout 检查是否已经退出登录
func CheckLogout(token string) bool {
	keys := global.Redis.Keys("logout_*").Val()
	if utils.InList("logout_"+token, keys) {
		//如果redis中存在，说明已经退出登录
		return true
	}
	//如果redis中不存在，说明没有退出登录
	return false
}
