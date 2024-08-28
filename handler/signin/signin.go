package signinHandler

import (
	"encoding/json"
	App "github.com/fossyy/WebAppTemplate/server"
	"github.com/fossyy/WebAppTemplate/session"
	"github.com/fossyy/WebAppTemplate/utils"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func POST(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		message := map[string]string{
			"error": "Invalid request payload",
		}
		json.NewEncoder(w).Encode(message)
		return
	}

	userData, _ := App.Server.Database.GetUser(req.Email)
	if req.Email == userData.Email && utils.CheckPasswordHash(req.Password, userData.Password) {
		token := session.MakeSession(userData.Username, userData.Email)
		message := map[string]string{
			"token": token,
		}
		json.NewEncoder(w).Encode(message)
		return
	}
	message := map[string]string{
		"error": "User atau password salah",
	}
	json.NewEncoder(w).Encode(message)
	return
}
