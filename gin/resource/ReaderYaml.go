package resource

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
)

var Database *database
var Server *server

type conf struct {
	Server server   `yaml:"server"`
	DB     database `yaml:"database"`
}

// server 服务器配置
type server struct {
	ApplicationName string `yaml:"applicationName"`
	Port            string `yaml:"port"`
}

// database 数据库配置
type database struct {
	Type            string `yaml:"type"`
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	UserName        string `yaml:"username"`
	Password        string `yaml:"password"`
	DbName          string `yaml:"dbname"`
	MaxIdleConn     int    `yaml:"maxIdleConn"`
	MaxOpenConn     int    `yaml:"maxOpenConn"`
	ConnMaxLifetime int    `yaml:"connMaxLifetime"`
}

func InitConf(dataFile string) {
	// 解决相对路经下获取不了配置文件问题
	_, filename, _, _ := runtime.Caller(0)
	filePath := path.Join(path.Dir(filename), dataFile)
	_, err := os.Stat(filePath)
	if err != nil {
		log.Printf("config file path %s not exist", filePath)
	}
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	c := new(conf)
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Printf("Unmarshal: %v", err)
	}
	log.Printf("load conf success\n %v", c)
	// 绑定到外部可以访问的变量中
	Server = &c.Server
	Database = &c.DB
}
