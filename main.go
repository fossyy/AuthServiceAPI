package main

import (
	"fmt"
	"github.com/fossyy/WebAppTemplate/db"
	"github.com/fossyy/WebAppTemplate/middleware"
	"github.com/fossyy/WebAppTemplate/routes"
	App "github.com/fossyy/WebAppTemplate/server"
	"github.com/fossyy/WebAppTemplate/utils"
)

func main() {
	Addr := fmt.Sprintf("%s:%s", utils.Getenv("SERVER_HOST"), utils.Getenv("SERVER_PORT"))

	dbUser := utils.Getenv("DB_USERNAME")
	dbPass := utils.Getenv("DB_PASSWORD")
	dbHost := utils.Getenv("DB_HOST")
	dbPort := utils.Getenv("DB_PORT")
	dbName := utils.Getenv("DB_NAME")

	database := db.NewMySQL(dbUser, dbPass, dbHost, dbPort, dbName)
	App.Server = App.NewServer(Addr, middleware.Handler(routes.Setup()), database)
	fmt.Printf("Listening on http://%s\n", App.Server.Addr)
	err := App.Server.ListenAndServe()
	if err != nil {
		return
	}
}
