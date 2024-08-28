package routes

import (
	signinHandler "github.com/fossyy/WebAppTemplate/handler/signin"
	signupHandler "github.com/fossyy/WebAppTemplate/handler/signup"
	userHandler "github.com/fossyy/WebAppTemplate/handler/user"
	"net/http"
)

func Setup() *http.ServeMux {
	handler := http.NewServeMux()

	handler.HandleFunc("POST /signin", func(w http.ResponseWriter, r *http.Request) {
		signinHandler.POST(w, r)
	})

	handler.HandleFunc("POST /signup", func(w http.ResponseWriter, r *http.Request) {
		signupHandler.POST(w, r)
	})

	handler.HandleFunc("GET /user", func(w http.ResponseWriter, r *http.Request) {
		userHandler.GET(w, r)
	})

	return handler
}
