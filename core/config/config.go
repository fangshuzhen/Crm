package config

import (
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var Conf Config

type Config struct {
	Server Server      `toml:"Server"`
	Mongo  MongoDb     `toml:"MongoDb"`
	MySQL  MongoDb     `toml:"MySQL"`
	Micro  MicroConfig `toml:"Micro"`
}

type Server struct {
	Host string
	Port int
	Name string
}

type MySQL struct {
	Url      string
	Name     string
	UserName string
}

type MongoDb struct {
	Url            []string
	Name           string
	UserName       string
	Password       string
	ReplicaSetName string
	Source         string
	Timeout        int64
	Type           string
	MaxPoolSize    int
}

type MicroConfig struct {
	Name      string   //服务名称
	Etcd      []string //ETCD集群地址
	Address   string   //当前地址
	Advertise string   //广播地址
}

//读取配置到Config
func InitConfig(defaultConfigPath string) error {
	absPath, err := filepath.Abs(defaultConfigPath)
	if err != nil {
		return err
	}

	if _, err := toml.DecodeFile(absPath, &Conf); err != nil {
		return err
	}

	return nil
}

//存储日志文件
const (
	Log_FILE_PATH = "E:/"
	LOG_FILE_NAME = "hello_"
)
