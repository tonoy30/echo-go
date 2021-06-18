package main

import (
	"github.com/tonoy30/echo-go/internal/api"
	"github.com/tonoy30/echo-go/pkg/data"
	"github.com/tonoy30/echo-go/pkg/settings"
)

func main() {
	s := settings.NewSettings()
	db := data.NewConnection(s)
	defer db.Disconnect()

	application := api.New(s, db.Client)
	application.Start()
}
