package config

import (
	"io/ioutil"

	"github.com/Sirlanri/distiot-user-server/server/log"
	"gopkg.in/yaml.v3"
)

type Conf struct {
	EmailToken string `yaml:"emailtoken"`
	Port       int    `yaml:"port"`
}

//全局配置文件
var Config Conf

func ReadYaml() (*Conf, error) {
	buf, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}
	var conf Conf
	err = yaml.Unmarshal(buf, &conf)
	if err != nil {
		log.Log.Errorln("配置文件读取失败", err.Error())
		return nil, err
	}

	return &conf, nil
}
