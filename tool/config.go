package tool

import (
	"encoding/json"
	"io/ioutil"
)

var Cfg *Config = nil

func GetConfig() *Config {
	// 用于给别的方法获取全局的配置对象使用
	return Cfg
}

type Config struct {
	AppName string      `json:"app_name"`
	AppMode string      `json:"app_mode"`
	AppHost string      `json:"app_host"`
	AppPort string      `json:"app_port"`
	Email   EmailConfig `json:"email"`
}

type EmailConfig struct {
	User string `json:"user"`
	Pass string `json:"pass"`
	Host string `json:"host"`
	Port string `json:"port"`
}

func ParseConfig(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(file, &Cfg)
	if err != nil {
		return nil, err
	}
	return Cfg, nil
}
