package routes

import (
	"anomaly_detector/internal/repository"
)

type PayloadMessage struct {
	Label    string `json:"label"`
	EntityID int32  `json:"id"`
	Value    int32  `json:"value"`
}

type StatRequestMessage struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Iterator int32  `json:"iterator"`
}

type ICollector interface {
	HandleEvent(entityID int32, eventLabel string, eventValue int32)
	GetState(from, to string, iterator int32) (repository.EventPreparedList, error)
}

type appInfo struct {
	OK        bool   `json:"ok"`
	BuildHash string `json:"build_hash"`
	BuildTime string `json:"build_time"`
	AppName   string `json:"app_name"`
}
