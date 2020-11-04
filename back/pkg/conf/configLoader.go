package conf

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

var (
	//ValidConfigFileExtension    = []string{"json", "ini", "yaml", "yml", "xml"}
	ValidConfigFileExtension    = []string{"yaml", "yml"}
	validConfigFileExtensionMap = make(map[string]bool)
	Conf                        Config
	v                           *viper.Viper
)

func init() {
	// 初始化支持的config文件类型
	for _, ext := range ValidConfigFileExtension {
		validConfigFileExtensionMap[ext] = true
	}
	// 将日志格式设置为json
	log.SetFormatter(&log.JSONFormatter{})
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
	v = GetViper()
	// viper进行设置 读配置文件
	v.SetConfigFile("./application.yml")
	v.AddConfigPath("../back/config")
	v.SetConfigName("application")
	v.SetConfigType("yml")
	err := v.ReadInConfig()
	if err != nil {
		log.Panicf("failed to parse config file [%s]", err.Error())
	}
	if err = v.Unmarshal(&Conf); err != nil {
		log.Panicf("failed to decode into struct [%s]", err.Error())
	}
}

// 判断是否是支持种类的config文件类型
func IsSupportConfigFileType(filename string) bool {
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(filename), "."))
	return validConfigFileExtensionMap[ext]
}

// ReadConfig 读取配置
func ReadConfig() {
	fmt.Println(Conf)
}

// NewViper 返回一个Viper实例
func GetViper() *viper.Viper {
	if v == nil {
		v = viper.New()
	}
	return v
}
