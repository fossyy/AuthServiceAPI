package main

import (
	"fmt"
	"github.com/fossyy/WebAppTemplate/middleware"
	"github.com/fossyy/WebAppTemplate/routes"
	"net/http"
)

func main() {
	serverAddr := "localhost:8000"
	server := http.Server{
		Addr:    serverAddr,
		Handler: middleware.Handler(routes.Setup()),
	}

	fmt.Printf("Listening on http://%s\n", serverAddr)
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
