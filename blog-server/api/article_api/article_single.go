package article_api

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

// 查询单篇文章

func (ArticleApi) ArticleSingleView(c *gin.Context) {
	id := c.Param("id")
	var articleModel models.ArticleModel
	err := global.DB.Take(&articleModel, id).Error
	if err != nil {
		res.FailWithMessage("文章不存在", c)
		return
	}
	res.OkWithData(articleModel, c)
}
