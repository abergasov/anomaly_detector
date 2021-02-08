package storage

import (
	"anomaly_detector/internal/config"
	"anomaly_detector/internal/logger"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // justifying it
	"github.com/jmoiron/sqlx"
)

type DBConnector struct {
	*sqlx.DB
}

func InitDBConnect(cnf *config.AppConfig) *DBConnector {
	conStr := fmt.Sprintf("%s:%s@(%s:%s)/%s", cnf.ConfigDB.User, cnf.ConfigDB.Pass, cnf.ConfigDB.Address, cnf.ConfigDB.Port, cnf.ConfigDB.DBName)
	db, err := sqlx.Connect("mysql", conStr)
	if err != nil {
		logger.Fatal(
			fmt.Sprintf("error connect to db %s %s@%s:%s", cnf.ConfigDB.DBName, cnf.ConfigDB.User, cnf.ConfigDB.Address, cnf.ConfigDB.Port),
			err,
		)
	}
	return &DBConnector{
		db,
	}
}
