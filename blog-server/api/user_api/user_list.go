package user_api

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/models/ctype"
	"server/models/res"
	"server/service/common"
	"server/utils/jwts"
	"strings"
)

func (UserApi) UserListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var page models.PageInfo
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var users []models.UserModel
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page,
		Debug:    false,
	})
	for _, user := range list {
		//判断权限
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			//如果是非管理员
			user.UserName = ""
		}
		//脱敏
		Desensitization(&user)
		users = append(users, user)
	}

	res.OkWithList(users, count, c)

}
func Desensitization(user *models.UserModel) {
	if len(user.Tel) == 11 {
		user.Tel = user.Tel[:3] + "****" + user.Tel[7:]
	} else {
		user.Tel = ""
	}
	eList := strings.Split(user.Email, "@")
	if len(eList) != 2 {
		user.Email = ""
		return
	}
	user.Email = eList[0][:1] + "****@" + eList[1]
}
