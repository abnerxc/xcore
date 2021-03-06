package xcore

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

//全局变量
var (
	G_APP_ENV  string //全局环境
	G_APP_PATH string //全局APP路径
	G_DB       *gorm.DB
	G_REDIS    *redis.Client
	G_CFG      *viper.Viper
	G_LOG      *logrus.Logger
)
