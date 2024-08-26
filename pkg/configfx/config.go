package configfx

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	AppName string
}

type Config struct {
	*viper.Viper
}

func New(p Params) *Config {
	envConf := fmt.Sprintf("config/%s/config.yaml", p.AppName)
	return &Config{getConfig(envConf)}
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
