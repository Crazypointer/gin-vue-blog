package global

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"server/config"
)

// 存放全局变量

var (
	Config         *config.Config
	DB             *gorm.DB
	Log            *logrus.Logger
	Router         *gin.Engine
	MysqlLog       logger.Interface
	ImageWhiteList = []string{
		".jpg",
		".png",
		".jpeg",
		".ico",
		".tiff",
		".gif",
		".svg",
		".webp",
	}
	Redis *redis.Client
)
