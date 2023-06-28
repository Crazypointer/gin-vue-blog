package article_api

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/models/res"
	"server/service/common"
)

func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr models.PageInfo //分页查询参数
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	list, count, _ := common.ComList(models.ArticleModel{}, common.Option{
		PageInfo: cr,
	})
	res.OkWithList(list, count, c)
}
