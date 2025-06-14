package base

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// 定义结构
type Config struct {
	Http struct {
		Port int
	}
	Db struct {
		Path string
	}
}

// 配置全局对象
var Conf *Config

// 初始化函数
func InitConfig(path string) {
	Conf = &Config{}

	// 读取配置文件
	fileData, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(fileData, &Conf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
