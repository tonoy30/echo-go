package main

import (
	"fmt"

	"github.com/tonoy30/echo-go/pkg/config"
	"github.com/tonoy30/echo-go/pkg/data"
)

func main() {
	settings := config.NewSettings()
	db := data.NewConnection(settings)
	defer db.Disconnect()

	fmt.Println("Tonoy")
}
