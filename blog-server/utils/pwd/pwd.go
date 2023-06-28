package pwd

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashPwd hash密码
func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// ComparePwd 比较密码 hashPwd:数据库中的密码 pwd:用户输入的密码
func ComparePwd(hashPwd, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
