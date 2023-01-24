package config

import (
	"fmt"
	rss_conf "goober/application/rss/config"

	"github.com/spf13/viper"
)

type MysqlConf struct {
	DataSourceName string `mapstructure:"dataSourceName"`
	MaxOpenConns   int    `mapstructure:"maxOpenConns"`
	MaxIdleConns   int    `mapstructure:"maxIdleConns"`
}

type DebugConf struct {
	Pprof bool `mapstructure:"pprof"`
}

type appConf struct {
	MySql MysqlConf          `mapstructure:"mysql"`
	Mode  string             `mapstructure:"mode"`
	Debug DebugConf          `mapstructure:"debug"`
	Rss   rss_conf.RssConfig `mapstructure:"rss"`
}

// 应用总配置
var AppConf = &appConf{}

func initConfig() {
	e := viper.Unmarshal(AppConf)

	rss_conf.SetConfig(&AppConf.Rss)

	if e != nil {
		fmt.Println("app conf unmarshal: error", e)
	}

}
