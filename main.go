package main

import (
	"fmt"

	"github.com/kswope/viper-experiment/config"
)

func main() {

	config.Setup()
	fmt.Println(config.Data)

}
