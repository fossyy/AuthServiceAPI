package session

import (
	"errors"
	"github.com/fossyy/WebAppTemplate/utils"
)

type Session struct {
	Username string
	Email    string
}

var sessions = make(map[string]*Session)

func GetSession(sessionID string) (Session, error) {
	userSession, ok := sessions[sessionID]
	if !ok {
		return Session{}, errors.New("session not found")
	}
	return *userSession, nil
}

func MakeSession(username string, email string) string {
	id := utils.GenerateRandomString(64)
	sessions[id] = &Session{
		Username: username,
		Email:    email,
	}
	return id
}
