package util

import "github.com/spf13/viper"

func InitConfig(configFile string) error {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	return viper.ReadInConfig()
}