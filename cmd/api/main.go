package main

import (
	"fmt"

	"github.com/tonoy30/echo-go/pkg/settings"
	"github.com/tonoy30/echo-go/pkg/data"
)

func main() {
	settings := settings.NewSettings()
	db := data.NewConnection(settings)
	defer db.Disconnect()

	fmt.Println("Tonoy")
}
