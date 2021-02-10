package gathering

import (
	"anomaly_detector/internal/config"
	"anomaly_detector/internal/routes"
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

type AppGatheringRouter struct {
	appBuildInfo   []byte
	FastHTTPEngine *router.Router
	config         *config.AppConfig
	collector      ICollector
}

func InitGatheringRouter(cnf *config.AppConfig, cm ICollector, appName, appHash, appBuild string) *AppGatheringRouter {
	b, _ := json.Marshal(routes.AppInfo{
		OK:        true,
		AppName:   appName,
		BuildHash: appHash,
		BuildTime: appBuild,
	})
	return &AppGatheringRouter{
		FastHTTPEngine: router.New(),
		config:         cnf,
		collector:      cm,
		appBuildInfo:   b,
	}
}

func (ar *AppGatheringRouter) InitRoutes() *router.Router {
	ar.FastHTTPEngine.GET("/", ar.Index)
	ar.FastHTTPEngine.POST("/gather", ar.Gather)
	ar.FastHTTPEngine.POST("/stat", ar.GetStat)
	return ar.FastHTTPEngine
}

func (ar *AppGatheringRouter) Index(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	ctx.SetStatusCode(http.StatusOK)
	_, _ = ctx.Write(ar.appBuildInfo)
}
