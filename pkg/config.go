package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
    auth_token	string `mapstructure:"AUTH_TOKEN"`
}

func LoadConfig() (auth_token string){
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	auth_token = fmt.Sprint(viper.Get("auth_token"))

	return auth_token
}