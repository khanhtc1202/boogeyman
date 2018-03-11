package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var config *Config

type Config struct {
	RankerConf RankerConf
}

func GetEnv() string {
	if production {
		return "production"
	}
	return "development"
}

func GetConfig() *Config {
	if config == nil {
		_, err := toml.Decode(*ReadSettingData(), &config)
		if err != nil {
			fmt.Println(fmt.Sprintf("err: %v", err))
			return nil
		}
	}

	return config
}

func ReadSettingData() *string {
	f, err := Asset("config." + GetEnv() + ".tml")
	if err != nil {
		fmt.Println(fmt.Sprintf("err %v", err))
		return nil
	}
	str := string(f)
	return &str
}
