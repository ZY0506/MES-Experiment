package mysql

import (
	"MES/exp1/settings"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var database *gorm.DB

// Init 初始化数据库
func Init(cfg *settings.MySQLConf) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.Timeout)
	var err error
	database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // 禁用默认事务
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // 禁用单数表名
			NoLowerCase:   false, // 禁用小写表名
		},
	})
	if err != nil {
		zap.L().Error("connect mysql failed", zap.Error(err))
		return
	}
	sqlDB, err := database.DB()
	if err != nil {
		zap.L().Error("get _mysql db failed", zap.Error(err))
		return
	}
	err = sqlDB.Ping()
	if err != nil {
		zap.L().Error("ping _mysql failed", zap.Error(err))
		return
	}
	// 设置连接池
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}
