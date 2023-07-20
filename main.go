package main

import (
	"fmt"
	"log"
	"main/config"
)

func main() {
	// テスト出力
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Println("test")
}
