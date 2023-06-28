package main

import (
	"fmt"
	"server/core"
	"server/global"
	"server/utils/jwts"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NikeName: "user",
		Role:     1,
		UserID:   1,
	})
	fmt.Println(token, err)
	claims, err := jwts.ParseToken(token)
	fmt.Println(claims, err)
}
