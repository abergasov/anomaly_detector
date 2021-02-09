package session

import (
	"anomaly_detector/internal/config"

	"github.com/labstack/echo/v4"
)

type Session struct {
	cnf *config.AppConfig
}

func InitSessionManager(cnf *config.AppConfig) *Session {
	return &Session{
		cnf: cnf,
	}
}

func (s *Session) AuthMiddleware(username, password string, c echo.Context) (bool, error) {
	if username == s.cnf.AuthUser && password == s.cnf.AuthToken {
		return true, nil
	}
	return false, nil
}
