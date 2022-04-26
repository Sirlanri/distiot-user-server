package config

import (
	"io/ioutil"

	"github.com/Sirlanri/distiot-user-server/server/log"
	"gopkg.in/yaml.v3"
)

func init() {
	ReadYaml()
}

type Conf struct {
	EmailToken string `yaml:"emailtoken"`
	HttpPort   string `yaml:"httpport"`
	RpcPort    string `yaml:"rpcport"`
	RedisAddr  string `yaml:"redisAddr"`
	RedisName  string `yaml:"redisName"`
	RedisPW    string `yaml:"redisPW"`
	RedisDB    int    `yaml:"redisDB"`
	MysqlUrl   string `yaml:"mysqlUrl"`
}

//全局配置文件
var Config Conf

func ReadYaml() {
	buf, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Log.Warnln("server-config ReadYaml 读取配置文件失败", err.Error())
		return
	}
	err = yaml.Unmarshal(buf, &Config)
	if err != nil {
		log.Log.Errorln("配置文件读取失败", err.Error())
	}

}
