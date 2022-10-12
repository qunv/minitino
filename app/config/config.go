package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App App
}

type App struct {
	RootName string
	Intro    string
	FindOn   FindOn
}

type FindOn struct {
	Github  string
	Twitter string
}

func LoadConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var c Config
	err = viper.Unmarshal(&c)

	if err != nil {
		panic(fmt.Errorf("fatal error unmarshal: %w", err))
	}
	return c
}
