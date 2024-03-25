package app

import (
	"lambda-func/api"
	"lambda-func/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() *App {
	dbStore := database.NewDynamoDBClient()
	apiHandler := api.NewApiHandler(dbStore)

	return &App{
		ApiHandler: *apiHandler,
	}
}
