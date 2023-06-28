package article_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

func (ArticleApi) ArticleRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var articleList []models.ArticleModel
	count := global.DB.Find(&articleList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("文章不存在", c)
		return
	}
	//TODO 如果当前标签下有文章，则不能删除
	global.DB.Delete(&articleList)
	res.OkWithMessage(fmt.Sprintf("共删除：%d 篇文章", count), c)
}
