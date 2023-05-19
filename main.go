package main

import (
	db "github.com/JY8752/go-onion-architecture-sample/infrastructure"
	"github.com/JY8752/go-onion-architecture-sample/registory"
	ui "github.com/JY8752/go-onion-architecture-sample/userinterface"
)

func main() {
	// db
	db := db.NewDBClient()
	defer db.Client.Close()

	// registory
	registory := registory.NewRegistory(db)

	// echo
	apiClient := ui.NewApiClient(registory)
	apiClient.RegisterRoute()
	apiClient.Start()
}
