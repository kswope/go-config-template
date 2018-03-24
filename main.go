package main

import (
	"fmt"

	"github.com/kswope/go-config-template/config"
)

func main() {

	config.Setup()
	fmt.Printf("%+v", config.Data)

}
