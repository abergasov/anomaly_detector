package routes

import (
	"anomaly_detector/internal/config"
	"encoding/json"
	"net/http"

	"github.com/valyala/fasthttp"

	"github.com/fasthttp/router"
)

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
	strOK              = []byte(`{"ok":true}`)
)

type AppRouter struct {
	appBuildInfo   []byte
	FastHTTPEngine *router.Router
	config         *config.AppConfig
	collector      ICollector
}

func InitRouter(cnf *config.AppConfig, cm ICollector, appName, appHash, appBuild string) *AppRouter {
	b, _ := json.Marshal(appInfo{
		OK:        true,
		AppName:   appName,
		BuildHash: appHash,
		BuildTime: appBuild,
	})
	return &AppRouter{
		FastHTTPEngine: router.New(),
		config:         cnf,
		collector:      cm,
		appBuildInfo:   b,
	}
}

func (ar *AppRouter) InitRoutes() *router.Router {
	ar.FastHTTPEngine.GET("/", ar.Index)
	ar.FastHTTPEngine.POST("/gather", ar.Gather)
	return ar.FastHTTPEngine
}

func (ar *AppRouter) Index(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	ctx.SetStatusCode(http.StatusOK)
	_, _ = ctx.Write(ar.appBuildInfo)
}
