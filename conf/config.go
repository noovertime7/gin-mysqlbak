package conf

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
)

var Config *ini.File

func init() {
	if Config != nil {
		return
	}
	path, err := os.Getwd()
	cfg, err := ini.Load(path + "/conf/config.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}
	Config = cfg
}

func GetStringConf(section, key string) string {
	return Config.Section(section).Key(key).MustString("获取string失败")
}

func GetBoolConf(section, key string) bool {
	return Config.Section(section).Key(key).MustBool(false)
}

func GetIntConf(section, key string) int {
	return Config.Section(section).Key(key).MustInt(0)
}
