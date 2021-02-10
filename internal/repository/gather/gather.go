package gather

import (
	"anomaly_detector/internal/logger"
	"anomaly_detector/internal/repository"
	"strings"
	"sync"
	"time"
)

var (
	// has iterator, data from/to
	sqlLimited = "SELECT event_date, entity_id, event_counter FROM events_count WHERE ec_id > ? AND event_date > ? AND event_date < ? ORDER BY ec_id ASC LIMIT 500"
	// has iterator, data from
	sqlUnLimited = "SELECT event_date, entity_id, event_counter FROM events_count WHERE ec_id > ? AND event_date > ? ORDER BY ec_id ASC LIMIT 500"
	// only start date have
	sqlStart = "SELECT event_date, entity_id, event_counter FROM events_count WHERE event_date > ? ORDER BY ec_id ASC LIMIT 500"
)

type DataGather struct {
	collector chan event
	dataMx    sync.Mutex
	data      map[int32]map[string]int32 // entityID -> eventLabel -> valuesSUM
	db        IStorageSaver
	pool      sync.Pool
}

func NewDataGather(d IStorageSaver) *DataGather {
	dg := &DataGather{
		collector: make(chan event, 1000),
		data:      make(map[int32]map[string]int32),
		db:        d,
		pool: sync.Pool{
			New: func() interface{} {
				return &event{}
			},
		},
	}

	// put data from channel to map
	go dg.collectEvents()
	// store data from map
	go dg.saveEvents()
	return dg
}

func (d *DataGather) collectEvents() {
	for i := range d.collector {
		d.dataMx.Lock()
		if _, ok := d.data[i.eID]; !ok {
			d.data[i.eID] = make(map[string]int32)
		}
		if _, ok := d.data[i.eID][i.label]; !ok {
			d.data[i.eID][i.label] = 0
		}
		d.data[i.eID][i.label] += i.val
		d.dataMx.Unlock()
	}
}

func (d *DataGather) saveEvents() {
	for range time.Tick(2 * time.Second) {
		d.dataMx.Lock()
		counterData := d.data
		d.data = make(map[int32]map[string]int32, len(counterData))
		d.dataMx.Unlock()
		values := make([]interface{}, 0, 30)
		placeHolders := make([]string, 0, 10)
		for j, v := range counterData {
			for l, c := range v {
				placeHolders = append(placeHolders, "(?,?,?)")
				values = append(values, j, l, c)
			}
			if len(placeHolders) >= 10 {
				d.insertData(placeHolders, values)
				values = make([]interface{}, 0, 30)
				placeHolders = make([]string, 0, 10)
			}
		}
		if len(placeHolders) > 0 {
			d.insertData(placeHolders, values)
		}
	}
}

func (d *DataGather) HandleEvent(entityID int32, eventLabel string, eventValue int32) {
	d.collector <- event{eID: entityID, label: eventLabel, val: eventValue}
}

func (d *DataGather) HandleEventWithPool(entityID int32, eventLabel string, eventValue int32) {
	s := d.pool.Get().(*event)
	s.eID = entityID
	s.label = eventLabel
	s.val = eventValue
	d.collector <- *s
	d.pool.Put(&s)
}

func (d *DataGather) GetState(data repository.StatRequestMessage) (repository.EventPreparedList, error) {
	var p []repository.EventPrepared
	var err error
	if data.To != "" && data.Iterator > 0 {
		err = d.db.Select(&p, sqlLimited, data.Iterator, data.From, data.To)
	} else if data.Iterator == 0 && data.To == "" {
		err = d.db.Select(&p, sqlStart, data.From)
	} else {
		err = d.db.Select(&p, sqlUnLimited, data.Iterator, data.From)
	}

	if err != nil {
		logger.Error("error load stat", err)
	}
	return p, err
}

func (d *DataGather) insertData(placeHolders []string, values []interface{}) {
	placeStr := strings.Join(placeHolders, ",")
	sqlI := "INSERT INTO events_count (entity_id,event_label,event_counter) VALUES " + placeStr
	_, err := d.db.Exec(sqlI, values...)
	if err != nil {
		logger.Error("error insert data", err) // todo mark data
	}
}
