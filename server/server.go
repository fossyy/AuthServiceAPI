package App

import (
	"github.com/fossyy/WebAppTemplate/db"
	"net/http"
)

var Server App

type App struct {
	http.Server
	Database db.Database
}

func NewServer(addr string, handler http.Handler, database db.Database) App {
	return App{
		Server: http.Server{
			Addr:    addr,
			Handler: handler,
		},
		Database: database,
	}
}
