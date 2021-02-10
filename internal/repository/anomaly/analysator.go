package anomaly

import (
	"anomaly_detector/internal/config"
	"anomaly_detector/internal/logger"
	"anomaly_detector/internal/repository"
	"bytes"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Analyser struct {
	gatherURL    string
	client       *http.Client
	aDataMU      sync.RWMutex
	analysedData []repository.EventAnalysed
}

func InitAnalyser(conf *config.AppConfig) *Analyser {
	a := &Analyser{
		gatherURL: conf.GatherURL + "/stat",
		client: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: 20,
			},
			Timeout: 10 * time.Second,
		},
	}
	go func() {
		var err error
		for range time.Tick(2 * time.Second) {
			err = a.catchInfo()
			if err != nil {
				logger.Error("error check data", err)
			}
		}
	}()
	return a
}

func (a *Analyser) GetCurrentState() []repository.EventAnalysed {
	a.aDataMU.RLock()
	defer a.aDataMU.RUnlock()
	return a.analysedData
}

func (a *Analyser) catchInfo() error {
	data, err := a.askGather(a.client)
	if err != nil {
		return err
	}
	events := &repository.EventPreparedList{}
	err = events.UnmarshalJSON(data)
	if err != nil {
		logger.Error("error parse message from gather", err)
		return err
	}
	a.analyseData(*events)
	return nil
}

func (a *Analyser) askGather(client *http.Client) ([]byte, error) {
	sReqest := &repository.StatRequestMessage{
		From:     time.Now().Format("2006-01-02 00:00:00"),
		Iterator: 1,
	}
	payload, _ := sReqest.MarshalJSON()

	req, err := http.NewRequest(http.MethodPost, a.gatherURL, bytes.NewBuffer(payload))
	if err != nil {
		logger.Error("error monitoring gather", err)
		return nil, err
	}

	response, err := client.Do(req)
	if err != nil {
		logger.Error("error send request to gather API", err)
		return nil, err
	}

	var body []byte
	body, err = ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		logger.Error("error parse API response", err)
		return nil, err
	}

	return body, nil
}

func (a *Analyser) analyseData(data repository.EventPreparedList) {
	aData := make([]repository.EventAnalysed, 0, len(data))
	for i := range data {
		aData = append(aData, repository.EventAnalysed(data[i]))
	}
	a.aDataMU.Lock()
	a.analysedData = aData
	a.aDataMU.Unlock()
}
