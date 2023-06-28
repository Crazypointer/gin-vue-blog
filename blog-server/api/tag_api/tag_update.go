package tag_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
)

// TagUpdateView 更新标签
// @Tags 标签管理
// @Summary 更新标签
// @Description 更新标签
// @Param data body TagRequest true "更新标签信息"
// @Router /api/tags/:id [put]
// @Success 200 {object} res.Response{data=string}
// @Produce json
func (TagApi) TagUpdateView(c *gin.Context) {
	var cr TagRequest
	id := c.Param("id")
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var tag models.TagModel
	err = global.DB.Take(&tag, id).Error
	if err != nil {
		res.FailWithMessage("标签不存在", c)
		return
	}

	// 结构体转map
	maps := structs.Map(&cr)
	//添加
	err = global.DB.Model(&tag).Updates(maps).Error
	if err != nil {
		global.Log.Error("修改标签失败", err)
		res.FailWithMessage("修改标签失败", c)
		return
	}
	res.OkWithMessage("修改标签成功", c)
}
