package article_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

func (ArticleApi) ArticleUpdateView(c *gin.Context) {
	var cr ArticleRequest
	id := c.Param("id")
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var article models.ArticleModel
	err = global.DB.Take(&article, id).Error
	if err == nil {
		res.FailWithMessage("文章不存在", c)
		return
	}

	// 结构体转map
	maps := structs.Map(&cr)
	//添加
	err = global.DB.Model(&article).Updates(maps).Error
	if err != nil {
		global.Log.Error("修改文章失败", err)
		res.FailWithMessage("修改文章失败", c)
		return
	}
	res.OkWithMessage("文章修改成功", c)
}
