package main

import (
  "fmt"
	"github.com/tonbi21/go-test/config"
)

func main() {
  fmt.Println("テスト")
	fmt.Println(config.Config.LogFile)
}