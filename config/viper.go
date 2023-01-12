package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// func GetMysqlConf() mysql.MysqlConf {
// 	viper.U
// }

func LoadConfig() {
	fmt.Println("loading config")
	viper.SetConfigName(".gbconf")       // name of config file (without extension)
	viper.SetConfigType("yaml")          // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/gb_blog/") // path to look for the config file in
	// viper.AddConfigPath("$HOME/.gb_blog") // call multiple times to add many search paths
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	initConfig()
}
