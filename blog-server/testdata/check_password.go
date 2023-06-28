package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "root" // 原始密码

	// 假设这是从数据库或其他存储中获取的哈希值
	//hashedPassword := "$2a$10$e7S/XnJFemf3LWxdEAcHDOCWYJnKP20EIRgOgzJ3uvddl0oqZieKu" // 1234
	hashedPassword := "$2a$10$QPs8GmlDOZe2aVZUOSY27e3a3b3yNIm1Bd4Jq19oo5r6xv97SVVdG" // root

	// 验证密码是否匹配
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println("密码不匹配:", err)
		return
	}

	fmt.Println("密码验证通过")
}
