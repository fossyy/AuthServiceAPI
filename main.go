package main

import (
	"fmt"
	"github.com/fossyy/WebAppTemplate/db"
	"github.com/fossyy/WebAppTemplate/routes"
	App "github.com/fossyy/WebAppTemplate/server"
)

func main() {
	database := db.NewMySQL("root", "Password123", "127.0.0.1", "3306", "tugashr")
	App.Server = App.NewServer("localhost:8000", routes.Setup(), database)
	fmt.Printf("Listening on http://%s\n", App.Server.Addr)
	err := App.Server.ListenAndServe()
	if err != nil {
		return
	}
}
