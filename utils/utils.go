package utils

import "github.com/spf13/viper"

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("local_config")
	return viper.ReadInConfig()
}