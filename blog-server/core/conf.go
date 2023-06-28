package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/fs"
	"io/ioutil"
	"log"
	"server/config"
	"server/global"
)

const ConfigFile = "settings.yaml"

// InitConf 读取yaml配置文件
func InitConf() {

	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf erro: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c) //使用 Go 语言中的 YAML 库将 YAML 格式的配置文件解析成 Go 语言结构体 存到c中
	if err != nil {
		log.Fatalf("config init Unmarshal: %v", err)
	}
	log.Println("config init success!")
	// 将读取到的数据封装到结构体中后 将结构体设置到全局变量中
	global.Config = c
}
func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Infof("config update success!")
	return nil
}
