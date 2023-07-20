package config

import (
	"log"
	// go-iniライブラリのインポート
	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

var Config ConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
	// config.iniファイルを読み込む
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	// 読み込んだ値を設定
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
