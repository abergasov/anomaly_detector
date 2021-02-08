package main

import (
	"anomaly_detector/internal/config"
	"anomaly_detector/internal/logger"
	"anomaly_detector/internal/repository/gather"
	"anomaly_detector/internal/routes"
	"anomaly_detector/internal/storage"
	"flag"
	"log"

	"github.com/valyala/fasthttp"

	"go.uber.org/zap"
)

var (
	appName   = "gathering_data"
	buildTime = "_dev"
	buildHash = "_dev"
	confFile  = flag.String("config", "./configs/common.yml", "Config file path")
)

func main() {
	flag.Parse()
	err := logger.NewLogger()
	if err != nil {
		log.Fatal("error log init", err)
	}
	logger.Info("start app")
	appConfig := config.InitConf(*confFile)

	logger.Info(
		"Try start server on port",
		zap.String("port", appConfig.AppPort),
		zap.String("url", "http://localhost:"+appConfig.AppPort),
	)

	dbConnect := storage.InitDBConnect(appConfig)
	dataGather := gather.NewDataGather(dbConnect)
	router := routes.InitRouter(appConfig, dataGather, appName, buildHash, buildTime)
	r := router.InitRoutes()
	err = fasthttp.ListenAndServe(":"+appConfig.AppPort, r.Handler)
	if err != nil {
		logger.Fatal("Router error", err)
	}
}
