package main

import (
	"MES/exp1/dao/mysql"
	"MES/exp1/logger"
	"MES/exp1/pkg/snowflake"
	"MES/exp1/routes"
	"MES/exp1/settings"
	"fmt"
)

func main() {
	// 读取配置文件
	settings.Init("config_yaml/config.yaml")

	// 初始化日志
	if err := logger.Init(&settings.Conf.LogConf, settings.Conf.SystemConf.Mode); err != nil {
		fmt.Printf("init logger failed,error: %v\n", err)
		return
	}

	// 初始化雪花算法
	snowflake.Init(1)

	// 初始化数据库
	mysql.Init(&settings.Conf.MySQLConf)

	routes.Run()

}
