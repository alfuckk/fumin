package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Result struct {
	fx.Out

	Config *viper.Config
}

func New(p Params) (Result, error) {

}
func New() *viper.Viper {
	envConf := os.Getenv("APP_CONF")
	if envConf == "" {
		flag.StringVar(&envConf, "conf", "config/local.yml", "config path, eg: -conf config/local.yml")
		flag.Parse()
	}
	if envConf == "" {
		envConf = "config/local.yml"
	}
	fmt.Println("load conf file:", envConf)
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
