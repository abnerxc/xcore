package global

import (
	"github.com/go-redis/redis"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

//全局变量
var (
	G_APP_ENV string
	G_DB      *gorm.DB
	G_REDIS   *redis.Client
	G_VP      *viper.Viper
	G_LOG     *logging.Logger
)
