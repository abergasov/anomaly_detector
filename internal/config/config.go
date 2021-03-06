package config

import (
	"anomaly_detector/internal/logger"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"

	"go.uber.org/zap"
)

type AppConfig struct {
	AppPort   string `yaml:"app_port"`
	AuthUser  string `yaml:"auth_user"`
	AuthToken string `yaml:"auth_token"`
	GatherURL string `yaml:"gather_url"`
	ConfigDB  DBConf `yaml:"conf_db"`
}

type DBConf struct {
	Address      string        `yaml:"address"`
	Port         string        `yaml:"port"`
	User         string        `yaml:"user"`
	Pass         string        `yaml:"pass"`
	DBName       string        `yaml:"db_name"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

func InitConf(confFile string) *AppConfig {
	logger.Info("Try read config from file", zap.String("path", confFile))

	file, errP := os.Open(filepath.Clean(confFile))
	if errP != nil {
		logger.Fatal("Error open file", errP, zap.String("file_path", confFile))
	}
	defer func() {
		e := file.Close()
		if e != nil {
			logger.Fatal("Error close config file", e)
		}
	}()

	var cfg AppConfig
	decoder := yaml.NewDecoder(file)
	errD := decoder.Decode(&cfg)
	if errD != nil {
		logger.Fatal("Invalid config file", errD, zap.String("file_path", confFile))
	}

	logger.Info("Config ok", zap.String("path", confFile))
	return &cfg
}
