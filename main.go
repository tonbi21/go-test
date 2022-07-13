package main

import (
	"fmt"

	"github.com/tonbi21/go-test/config"
)

func main() {
  fmt.Println("テスト")
	fmt.Print(config.Config.Port)
	fmt.Print(config.Config.SQLDriver)
	fmt.Print(config.Config.DbName)
	fmt.Print(config.Config.LogFile)

	config.LoadConfig()
}