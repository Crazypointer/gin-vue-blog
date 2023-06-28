package tag_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

// TagRemoveView 删除标签
// @Tags 标签管理
// @Summary 删除标签
// @Description 删除标签
// @Param data body models.RemoveRequest true "标签id列表"
// @Router /api/tags [delete]
// @Success 200 {object} res.Response{data=string}
// @Produce json
func (TagApi) TagRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var tagList []models.TagModel
	count := global.DB.Find(&tagList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("标签不存在", c)
		return
	}
	//TODO 如果当前标签下有文章，则不能删除
	global.DB.Delete(&tagList)
	res.OkWithMessage(fmt.Sprintf("共删除：%d 个标签", count), c)
}
