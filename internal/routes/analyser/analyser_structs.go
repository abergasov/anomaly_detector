package analyser

import (
	"anomaly_detector/internal/repository"

	"github.com/labstack/echo/v4"
)

type Resp struct {
	OK   bool                       `json:"ok"`
	Data []repository.EventAnalysed `json:"data"`
}

//go:generate mockgen -source=analyser_structs.go -destination=analyser_structs_mock.go -package=analyser
type IAnalyser interface {
	GetCurrentState() []repository.EventAnalysed
}

//go:generate mockgen -source=analyser_structs.go -destination=analyser_structs_mock.go -package=analyser
type ISessionManager interface {
	AuthMiddleware(username, password string, c echo.Context) (bool, error)
}
