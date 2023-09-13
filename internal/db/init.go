package db

import (
	"github.com/huermiaowu/miao-user/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"time"
)

var DB *gorm.DB

func Init(c *config.Config) error {
	var err error
	DB, err = gorm.Open(mysql.Open(c.Mysql.DSN))
	if err != nil {
		return err
	}
	if c.MysqlSlave.HasConfig() {
		err = DB.Use(
			dbresolver.Register(dbresolver.Config{
				Replicas:          []gorm.Dialector{mysql.Open(c.MysqlSlave.DSN)},
				Policy:            dbresolver.RandomPolicy{},
				TraceResolverMode: true,
			}).
				SetMaxIdleConns(10).
				SetConnMaxLifetime(time.Hour).
				SetMaxOpenConns(200))
		if err != nil {
			return err
		}
	}

	// 根据结构自动建表
	err = DB.AutoMigrate(&User{})
	if err != nil {
		return err
	}

	return nil
}
