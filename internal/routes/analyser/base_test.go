package analyser

import (
	"anomaly_detector/internal/config"
	"anomaly_detector/internal/logger"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"

	"github.com/golang/mock/gomock"
)

var session *MockISessionManager
var analyser *MockIAnalyser

func TestGetState(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cnf := &config.AppConfig{
		AppPort:   "1231",
		AuthToken: "abc",
		AuthUser:  "cdf",
	}
	engine := createRouter(ctrl, cnf)

	url := "/api/v1/state"
	method := http.MethodGet

	checkAuthWork(method, url, engine, t)

	println("check GetCurrentState valid")
	session.EXPECT().AuthMiddleware(gomock.Eq(cnf.AuthUser), gomock.Eq(cnf.AuthToken), gomock.Any()).Return(true, nil)
	analyser.EXPECT().GetCurrentState()
	w := performRequest(engine, method, url, cnf.AuthUser, cnf.AuthToken, nil)

	if w.Code != http.StatusOK {
		t.Errorf("expected %d code, got %d", http.StatusOK, w.Code)
	}
}

func createRouter(ctrl *gomock.Controller, cnf *config.AppConfig) *echo.Echo {
	_ = logger.NewLogger()
	session = NewMockISessionManager(ctrl)
	analyser = NewMockIAnalyser(ctrl)
	router := InitAnalyserRouter(cnf, analyser, session, "test", "test", "hash")
	return router.InitRoutes()
}

func checkAuthWork(method, url string, engine http.Handler, t *testing.T) {
	println(fmt.Sprintf("try check auth %s:%s with invalid auth", method, url))
	session.EXPECT().AuthMiddleware(gomock.Eq("d"), gomock.Eq("g"), gomock.Any())
	w := performRequest(engine, method, url, "d", "g", nil)
	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected %d code, got %d", http.StatusUnauthorized, w.Code)
	}
	println("end check auth")
}

func performRequest(r http.Handler, method, path, login, pass string, payload interface{}) *httptest.ResponseRecorder {
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	if payload != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	req.SetBasicAuth(login, pass)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
