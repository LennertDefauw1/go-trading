package main

import (
	"fmt"

	"github.com/LennertDefauw1/go-trading/src/config"
)

func main() {

	fmt.Println("Hello, World!")

	config, _ := config.LoadConfig("./config")

	fmt.Println(config.PrivateApiKey)
}
