package jwts

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"server/global"
	"time"
)

type JwtPayLoad struct {
	NikeName string `json:"nike_name"` // 昵称
	Role     int    `json:"role"`      // 权限  1 管理员  2 普通用户  3 游客
	UserID   int    `json:"user_id"`   // 用户id
}
type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

// GenToken 生成token

func GenToken(user JwtPayLoad) (string, error) {
	var MySecret = []byte(global.Config.Jwt.Secret)
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))), // 过期时间2小时
			Issuer:    global.Config.Jwt.Issuer,                                                     // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

// ParseToken 解析token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	var MySecret = []byte(global.Config.Jwt.Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		global.Log.Error(fmt.Sprintf("解析token失败：%s", err))
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("非法token")
}
