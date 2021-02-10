package repository

//easyjson:json
type EventPreparedList []EventPrepared

type EventPrepared struct {
	EventDate    string `json:"event_date" db:"event_date"`
	EntityID     int32  `json:"entity_id" db:"entity_id"`
	EventCounter int32  `json:"event_counter" db:"event_counter"`
}

type EventAnalysed struct {
	EventDate    string `json:"event_date"`
	EntityID     int32  `json:"entity_id"`
	EventCounter int32  `json:"event_counter"`
}

type StatRequestMessage struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Iterator int32  `json:"iterator"`
}
