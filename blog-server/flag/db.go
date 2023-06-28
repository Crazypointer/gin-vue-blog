package flag

import (
	"server/global"
	"server/models"
)

func Makemigrations() {
	global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	// 生成四张表的表结果
	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.BannerModel{},
		&models.TagModel{},
		&models.MessageModel{},
		&models.AdvertModel{},
		&models.UserModel{},
		&models.CommentModel{},
		&models.ArticleModel{},
		&models.MenuModel{},
		&models.MenuBannerModel{},
		&models.FeedBackModel{},
		&models.LoginDataModel{},
	)
	if err != nil {
		global.Log.Error("生成数据库表结构失败")
		return
	}
	global.Log.Infof("生成数据库表结构成功！")
}
