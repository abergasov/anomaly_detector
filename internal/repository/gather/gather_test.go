package gather

import (
	"anomaly_detector/internal/logger"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func BenchmarkSingleMap(b *testing.B) {
	_ = logger.NewLogger()
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()
	dbConnect := NewMockIStorageSaver(ctrl)
	dbConnect.EXPECT().Exec(gomock.Any(), gomock.Any()).MinTimes(1)

	collectorSingleMap := NewDataGather(dbConnect)
	for i := 0; i < b.N; i++ {
		collectorSingleMap.HandleEvent(12, "sample", 123)
	}
	time.Sleep(3 * time.Second)
}

func BenchmarkSingleMapWithPool(b *testing.B) {
	_ = logger.NewLogger()
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()
	dbConnect := NewMockIStorageSaver(ctrl)
	dbConnect.EXPECT().Exec(gomock.Any(), gomock.Any()).MinTimes(1)

	collectorSingleMap := NewDataGather(dbConnect)
	for i := 0; i < b.N; i++ {
		collectorSingleMap.HandleEventWithPool(12, "sample", 123)
	}
	time.Sleep(3 * time.Second)
}
