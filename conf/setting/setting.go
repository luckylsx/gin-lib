package setting

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"gopkg.in/yaml.v2"
)

// Setting 相应设置配置
type Setting struct {
	RunMode             string   `yaml:"runMode"`
	Server              server   `yaml:"server"`
	DspPoolRead         DataBase `yaml:"dsp-pool-read"`
	DspPoolWrite        DataBase `yaml:"dsp-pool-write"`
	DspBusinessPoolRead DataBase `yaml:"dsp-business-pool-read"`
	DspRecordRead       DataBase `yaml:"dsp-pool-record-read"`
	DspRecordWrite      DataBase `yaml:"dsp-pool-record-write"`
	Log                 Log      `yaml:"log"`
}

//服务配置
type server struct {
	HTTPPort     string        `yaml:"HTTPPort"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
}

// DataBase 数据库配置
type DataBase struct {
	Type     string `yaml:"type"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
}

// Log log
type Log struct {
	Path      string `yaml:"path"`
	Prefix    string `yaml:"prefix"`
	SQLPrefix string `yaml:"SQLPrefix"`
}

var conf = &Setting{}

func init() {
	env := GetEnv()
	yamlFile, err := ioutil.ReadFile("conf/env/." + env + ".yaml")
	if err != nil {
		panic(err)
	}

	if env == "pro" {
		gin.SetMode(gin.ReleaseMode)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		panic(err)
	}
}

// Conf 获取配置  外部调用使用
func Conf() *Setting {
	return conf
}

// GetEnv return the env
func GetEnv() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	return env
}
