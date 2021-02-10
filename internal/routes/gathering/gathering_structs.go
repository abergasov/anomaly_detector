package gathering

import (
	"anomaly_detector/internal/repository"
)

type PayloadMessage struct {
	Label    string `json:"label"`
	EntityID int32  `json:"id"`
	Value    int32  `json:"value"`
}

//go:generate mockgen -source=gathering_structs.go -destination=gathering_structs_mock.go -package=gathering
type ICollector interface {
	HandleEvent(entityID int32, eventLabel string, eventValue int32)
	GetState(data repository.StatRequestMessage) (repository.EventPreparedList, error)
}
