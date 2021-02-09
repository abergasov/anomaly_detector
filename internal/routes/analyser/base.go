package analyser

import (
	"anomaly_detector/internal/config"
	"anomaly_detector/internal/routes"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

type AppAnalyserRouter struct {
	appBuildInfo string
	EchoEngine   *echo.Echo
	config       *config.AppConfig
	analyser     IAnalyser
	session      ISessionManager
}

func InitAnalyserRouter(cnf *config.AppConfig, a IAnalyser, s ISessionManager, appName, appHash, appBuild string) *AppAnalyserRouter {
	b, _ := json.Marshal(routes.AppInfo{
		OK:        true,
		AppName:   appName,
		BuildHash: appHash,
		BuildTime: appBuild,
	})
	return &AppAnalyserRouter{
		EchoEngine:   echo.New(),
		config:       cnf,
		analyser:     a,
		session:      s,
		appBuildInfo: string(b),
	}
}

func (ar *AppAnalyserRouter) InitRoutes() *echo.Echo {
	ar.EchoEngine.GET("/", ar.Ping)
	groupV1 := ar.EchoEngine.Group("/api/v1", middleware.BasicAuth(ar.session.AuthMiddleware))
	groupV1.GET("/state", ar.GetState)
	return ar.EchoEngine
}

func (ar *AppAnalyserRouter) Ping(c echo.Context) error {
	return c.String(http.StatusOK, ar.appBuildInfo)
}

func (ar *AppAnalyserRouter) GetState(c echo.Context) error {
	return c.JSON(http.StatusOK, Resp{
		OK:   true,
		Data: ar.analyser.GetCurrentState(),
	})
}
