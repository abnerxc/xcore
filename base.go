package xcore

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

//启动引导
func Bootstrap(env string) {
	if env == "" {
		G_APP_ENV = "dev"
	} else {
		G_APP_ENV = env
	}
	//初始化配置读取
	initConfig()
	//初始化数据库配置
	if G_VP.IsSet("datasource") {
		initDB()
	}
	//初始化redis连接
	if G_VP.IsSet("redis") {
		initRedis()
	}
}

//资源关闭
func CloseRes() {
	_ = G_REDIS.Close()
}

//初始化配置文件
func initConfig() {
	fmt.Println("init config .....")
	dir, err := os.Getwd()
	if err != nil {
		panic("获取路径错误")
	}
	//读取yaml文件
	v := viper.New()
	//设置读取的配置文件
	v.SetConfigName(G_APP_ENV)
	//go,bin运行的路径
	v.AddConfigPath(filepath.FromSlash(dir + "/config/"))
	//设置配置文件类型
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	G_VP = v
}

//初始化数据库配置
func initDB() {
	fmt.Println("init db .....")
	G_DB = NewDBClient()
}

//初始化数据库配置
func initRedis() {
	fmt.Println("init redis .....")
	G_REDIS = NewRedis()
}