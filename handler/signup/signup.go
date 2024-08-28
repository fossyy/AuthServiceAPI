package signupHandler

import (
	"encoding/json"
	"github.com/fossyy/WebAppTemplate/db"
	"github.com/fossyy/WebAppTemplate/utils"
	"github.com/google/uuid"
	"net/http"
)

type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func POST(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req SignupRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		message := map[string]string{
			"error": "Invalid request payload",
		}
		json.NewEncoder(w).Encode(message)
		return
	}

	if !utils.IsValidEmail(req.Email) {
		message := map[string]string{
			"error": "Invalid email address",
		}
		json.NewEncoder(w).Encode(message)
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)

	newUser := db.User{
		UserID:   uuid.New(),
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	err = db.DB.Create(&newUser).Error

	if err != nil {
		message := map[string]string{
			"error": "Username atau Email sudah terdaftar",
		}
		json.NewEncoder(w).Encode(message)
		return
	}

	message := map[string]string{
		"message": "Berhasil register",
	}
	json.NewEncoder(w).Encode(message)
	return
}
