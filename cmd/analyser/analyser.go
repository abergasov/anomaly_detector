package main

import (
	"anomaly_detector/internal/config"
	"anomaly_detector/internal/logger"
	"anomaly_detector/internal/repository/anomaly_analysator"
	"anomaly_detector/internal/repository/session"
	"anomaly_detector/internal/routes/analyser"
	"flag"
	"log"

	"go.uber.org/zap"
)

var (
	appName   = "analyser_data"
	buildTime = "_dev"
	buildHash = "_dev"
	confFile  = flag.String("config", "./configs/common_analyser.yml", "Config file path")
)

func main() {
	flag.Parse()
	err := logger.NewLogger()
	if err != nil {
		log.Fatal("error log init", err)
	}
	logger.Info("start app")
	appConfig := config.InitConf(*confFile)
	sessionRepo := session.InitSessionManager(appConfig)
	analyserDetector := anomaly_analysator.InitAnalyser()
	logger.Info(
		"Try start analysing server on port",
		zap.String("port", appConfig.AppPort),
		zap.String("url", "http://localhost:"+appConfig.AppPort),
	)
	router := analyser.InitAnalyserRouter(appConfig, analyserDetector, sessionRepo, appName, buildHash, buildTime)
	err = router.InitRoutes().Start(":" + appConfig.AppPort)
	if err != nil {
		logger.Fatal("Router error", err)
	}
}
