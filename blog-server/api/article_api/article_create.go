package article_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/ctype"
	"server/models/res"
)

// 文章创建请求参数
/*
{
    "title": "abc"
	"keyword": "abc"
	"abstract": "abc"
	"content": "abc"
	"user_id": 1
	"user_nick_name": "abc"
	"user_avatar": "abc"
	"category": "abc"
	"source": "abc"
	"link": "abc"
	"banner_id": 1
	"banner_url": "abc"
	"tags": "abc"
}

*/

type ArticleRequest struct {
	CreatedAt string `json:"created_at" structs:"created_at"`      // 创建时间
	UpdatedAt string `json:"updated_at" structs:"updated_at"`      // 更新时间
	Title     string `json:"title" structs:"title"`                // 文章标题
	Keyword   string `json:"keyword,omit(list)" structs:"keyword"` // 关键字
	Abstract  string `json:"abstract" structs:"abstract"`          // 文章简介
	Content   string `json:"content,omit(list)" structs:"content"` // 文章内容

	LookCount     int `json:"look_count" structs:"look_count"`         // 浏览量
	CommentCount  int `json:"comment_count" structs:"comment_count"`   // 评论量
	DiggCount     int `json:"digg_count" structs:"digg_count"`         // 点赞量
	CollectsCount int `json:"collects_count" structs:"collects_count"` // 收藏量

	UserID       uint   `json:"user_id" structs:"user_id"`               // 用户id
	UserNickName string `json:"user_nick_name" structs:"user_nick_name"` //用户昵称
	UserAvatar   string `json:"user_avatar" structs:"user_avatar"`       // 用户头像

	Category string `json:"category" structs:"category"`        // 文章分类
	Source   string `json:"source,omit(list)" structs:"source"` // 文章来源
	Link     string `json:"link,omit(list)" structs:"link"`     // 原文链接

	BannerID  uint   `json:"banner_id" structs:"banner_id"`   // 文章封面图片id
	BannerUrl string `json:"banner_url" structs:"banner_url"` // 文章封面

	Tags ctype.Array `json:"tags" structs:"tags"` // 文章标签
}

func (ArticleApi) ArticleCreateView(c *gin.Context) {
	var cr ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	fmt.Printf("%+v", cr)

	var article models.ArticleModel
	// 通过标题查找
	err = global.DB.Take(&article, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("该标题已存在", c)
		return
	}
	//添加
	err = global.DB.Create(&models.ArticleModel{
		Title:        cr.Title,
		Keyword:      cr.Keyword,
		Abstract:     cr.Abstract,
		Content:      cr.Content,
		UserID:       cr.UserID,
		UserNickName: cr.UserNickName,
		UserAvatar:   cr.UserAvatar,
		Category:     cr.Category,
		Source:       cr.Source,
		Link:         cr.Link,
		BannerID:     cr.BannerID,
		BannerUrl:    cr.BannerUrl,
		Tags:         cr.Tags,
	}).Error
	if err != nil {
		global.Log.Error("添加文章失败", err)
		res.FailWithMessage("添加文章失败", c)
		return
	}
	res.OkWithMessage("添加文章成功", c)
}
