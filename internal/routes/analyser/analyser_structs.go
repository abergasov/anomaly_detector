package analyser

import "github.com/labstack/echo/v4"

type IAnalyser interface {
	GetCurrentState() []string
}

type Resp struct {
	OK   bool     `json:"ok"`
	Data []string `json:"data"`
}

type ISessionManager interface {
	AuthMiddleware(username, password string, c echo.Context) (bool, error)
}
