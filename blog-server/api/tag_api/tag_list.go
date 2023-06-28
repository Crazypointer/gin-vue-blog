package tag_api

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"server/models/res"
	"server/service/common"
)

// TagListView 标签列表
// @Tags 标签管理
// @Summary 标签列表
// @Description 标签列表
// @Param data query models.PageInfo false "查询参数"
// @Router /api/tags [get]
// @Success 200 {object} res.Response{data=res.ListResponse[models.TagModel]}
// @Produce json
func (TagApi) TagListView(c *gin.Context) {
	var cr models.PageInfo //分页查询参数
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	list, count, _ := common.ComList(models.TagModel{}, common.Option{
		PageInfo: cr,
	})
	// TODO: 需要展示标签下的文章数量
	res.OkWithList(list, count, c)
}
