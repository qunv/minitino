package config

import (
	"fmt"
	"github.com/qunv/minitino/app/models"
	"github.com/spf13/viper"
)

func LoadConfig() models.Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var c models.Config
	err = viper.Unmarshal(&c)

	if err != nil {
		panic(fmt.Errorf("fatal error unmarshal: %w", err))
	}
	return c
}
