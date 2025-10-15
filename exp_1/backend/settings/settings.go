package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	SystemConf `mapstructure:"system"`
	LogConf    `mapstructure:"log"`
	MySQLConf  `mapstructure:"mysql"`
}

type SystemConf struct {
	Mode string `mapstructure:"mode"`
	Port int    `mapstructure:"port"`
}

// MySQLConf 配置信息
type MySQLConf struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	Timeout      string `mapstruct:"timeout"`
	MaxOpenConns int    `mapstructure:"max_open_conns"` // _mysql 最大连接数
	MaxIdleConns int    `mapstructure:"max_idle_conns"` // _mysql 最大空闲连接数
}

// LogConf 配置信息
type LogConf struct {
	Level      string `mapstructure:"level"`
	Path       string `mapstructure:"path"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"` // 最多保留的备份文件数量
}

// Init 初始化配置
func Init(file string) {
	// 设置配置文件路径
	viper.SetConfigFile(file)
	// 监听配置文件变化
	viper.WatchConfig()
	// 配置文件发生变更时的回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件已更新:", e.Name)
		// 重新解析配置到结构体中
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("配置重新加载失败: %v\n", err))
		}
	})

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败: %v", err))
	}

	// 将配置解析到 Conf 结构体中
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("解析配置失败: %v", err))
	}
	return
}
