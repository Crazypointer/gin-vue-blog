package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/global"
	"server/models"
	"server/models/res"
)

//  TODO: 这个接口待测试

func (UserApi) UserRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var userList []models.UserModel
	count := global.DB.Find(&userList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("用户不存在", c)
		return
	}
	//开启事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// TODO: 删除用户后，与用户相关联的表也要删除 例如：用户收藏的文章，用户的评论等

		err = global.DB.Delete(&userList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除用户失败", c)
		return
	}

	res.OkWithMessage(fmt.Sprintf("共删除：%d 个用户", count), c)
}
