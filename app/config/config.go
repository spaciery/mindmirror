package config

import (
	"os"

	"github.com/spf13/viper"
)

type Cfg struct {
}

func (c *Cfg) loadConfig() {

	configpath, err := os.UserConfigDir()
	if err != nil {

	}

	viper.AddConfigPath(configpath + "/mindmirror/")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	err = viper.ReadInConfig()
	if err != nil {

	}

	err = viper.Unmarshal(&c)
	if err != nil {

	}

}
