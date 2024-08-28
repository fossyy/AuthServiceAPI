package userHandler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/fossyy/WebAppTemplate/session"
)

func GET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("Authorization")
	if token == "" {
		message := map[string]string{
			"error": "token required",
		}
		json.NewEncoder(w).Encode(message)
		return
	}

	splitToken := strings.Split(token, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		message := map[string]string{
			"error": "Invalid authorization format",
		}
		json.NewEncoder(w).Encode(message)
		return
	}

	userSession, err := session.GetSession(splitToken[1])
	if err != nil {
		message := map[string]string{
			"error": err.Error(),
		}
		json.NewEncoder(w).Encode(message)
		return
	}

	json.NewEncoder(w).Encode(userSession)
	return
}
