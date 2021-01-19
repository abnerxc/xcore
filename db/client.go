package db

import (
	"github.com/abnerxc/xcore/global"
	"github.com/mitchellh/mapstructure"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"log"
	"os"
	"time"
)

//数据库配置
type Dconfig struct {
	Type         string
	Dsn          map[string]string
	Maxidleconns int //设置空闲连接池中连接的最大数量
	Maxopenconns int //设置打开数据库连接的最大数量
	Maxidletime  int //设置连接可复用的最大时间
	Maxlifetime  int //设置连接连接的最大时间
}

var dconfig *Dconfig

//设置链接的数据源
func NewDBClient() *gorm.DB {
	dataSource := global.G_VP.GetStringMap("datasource")
	_ = mapstructure.Decode(dataSource["default"], &dconfig)
	db, err := gorm.Open(getDriver(dconfig.Type, dconfig.Dsn["slave"]), &gorm.Config{})
	if err != nil {
		panic("open data error:" + err.Error())
	}
	_ = db.Use(setConnectSource(dataSource))
	return db
}

func getDriver(dtype string, dsn string) gorm.Dialector {
	switch dtype {
	case "postgres":
		return postgres.Open(dsn)
	case "sqlite":
		return sqlite.Open(dsn)
	default: //mysql
		return mysql.Open(dsn)
	}
}

//数据库连接源配置
func setConnectSource(dataSource map[string]interface{}) *dbresolver.DBResolver {
	dbPlus := dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{getDriver(dconfig.Type, dconfig.Dsn["master"])},
		Replicas: []gorm.Dialector{getDriver(dconfig.Type, dconfig.Dsn["slave"])},
		Policy:   dbresolver.RandomPolicy{},
	}).SetConnMaxIdleTime(time.Duration(dconfig.Maxidletime) * time.Hour).
		SetConnMaxLifetime(time.Duration(dconfig.Maxlifetime) * time.Hour).
		SetMaxIdleConns(dconfig.Maxidleconns).
		SetMaxOpenConns(dconfig.Maxopenconns)
	if _, ok := dataSource["activity"]; ok {
		_ = mapstructure.Decode(dataSource["activity"], &dconfig)
		dbPlus.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{getDriver(dconfig.Type, dconfig.Dsn["master"])},
			Replicas: []gorm.Dialector{getDriver(dconfig.Type, dconfig.Dsn["slave"])},
			Policy:   dbresolver.RandomPolicy{},
		}, "activity").SetConnMaxIdleTime(time.Duration(dconfig.Maxidletime) * time.Hour).
			SetConnMaxLifetime(time.Duration(dconfig.Maxlifetime) * time.Hour).
			SetMaxIdleConns(dconfig.Maxidleconns).
			SetMaxOpenConns(dconfig.Maxopenconns)
	}
	return dbPlus
}

func setLogger() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 2 * time.Second, // Slow SQL threshold
			LogLevel:      logger.Silent,   // Log level
			Colorful:      false,           // Disable color
		},
	)
	return newLogger
}
