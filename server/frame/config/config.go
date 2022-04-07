package config

import (
	"errors"
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
)

// 配置信息
var Configuration *Config

// 启动配置参数
type CmdConfigParam struct {
	ConfigFilePath string // -e  配置文件路径
}

// 启动可执行文件时的参数
var startConfigParam *CmdConfigParam

type Config struct {
	App    *App    `yaml:"app"`
	Server *Server `yaml:"server"`
	Logger *Log    `yaml:"log"`
	Database *Database `yaml:"database"`
}

// 初始化函数将配置文件写入Config里
func init() {
	configFilePath := flag.String("c", "./config.yml", "配置文件路径，默认为可执行文件目录")
	flag.Parse()
	// 获取启动参数中，配置文件的绝对路径
	path, _ := filepath.Abs(*configFilePath)
	startConfigParam = &CmdConfigParam{ConfigFilePath: path}
	// 读取配置文件信息
	yc := &Config{}
	if err := LoadYml(startConfigParam.ConfigFilePath, yc); err != nil {
		panic(fmt.Sprintf("读取配置文件[%s]失败: %s", startConfigParam.ConfigFilePath, err.Error()))
	}
	Configuration = yc
}

// yaml读取工具函数
// 从指定路径加载yaml文件
func LoadYml(path string, out interface{}) error {
	yamlFileBytes, readErr := ioutil.ReadFile(path)
	if readErr != nil {
		return readErr
	}
	// yaml解析
	err := yaml.Unmarshal(yamlFileBytes, out)
	if err != nil {
		return errors.New("无法解析 [" + path + "] -- " + err.Error())
	}
	return nil
}

func LoadYmlByString(yamlStr string, out interface{}) error {
	// yaml解析
	return yaml.Unmarshal([]byte(yamlStr), out)
}

//==========App 配置信息=============

type App struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

func (a *App) GetAppInfo() string {
	return fmt.Sprintf("[%s:%s]", a.Name, a.Version)
}

//=========Server配置信息============
type Server struct {
	Address string `yaml:"address"`
}

//===========日志配置=================
type Log struct {
	Path   string `yaml:"path"`
	Level  uint32 `yaml:"level"`
	Stdout bool   `yaml:"stdout"`
}

//===========数据库==================
type Database struct {
	Driver string `yaml:"driver"`
	DbUrl string `yaml:"url"`
	Debug bool   `yaml:"debug"`

}
