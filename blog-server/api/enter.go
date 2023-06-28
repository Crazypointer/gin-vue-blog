package api

import (
	"server/api/advert_api"
	"server/api/article_api"
	"server/api/images_api"
	"server/api/menu_api"
	"server/api/settings_api"
	"server/api/tag_api"
	"server/api/user_api"
)

// 所有的接口都需要在这里绑定到结构体中

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
	MenuApi     menu_api.MenuApi
	UserApi     user_api.UserApi
	TagApi      tag_api.TagApi
	ArticleApi  article_api.ArticleApi
}

// 实例化一个对象 这个对象在router中使用

var ApiGroupApp = new(ApiGroup)
