package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type HttpConfig struct {
	Port string `yaml:"port"`
}

type Config struct {
	Http HttpConfig `yaml:"http"`
}

func New() (*Config, error) {
	viper.AddConfigPath("config/novel")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	fmt.Println(22)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
