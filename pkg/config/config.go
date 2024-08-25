package config

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	AppName string
}

func New(p Params) *viper.Viper {
	envConf := fmt.Sprintf("config/%s/config.yaml", p.AppName)

	return getConfig(envConf)
}

func getConfig(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
