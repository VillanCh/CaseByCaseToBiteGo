package demo

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestPipeline(t *testing.T) {
	// 生产者
	origin := make(chan *DemoData)
	go func() {
		defer close(origin)

		for i := 0; i < 10; i++ {
			origin <- &DemoData{
				Type: "origin type",
				Data: i,
			}
		}
	}()

	count := 0
	Pipeline1(origin, Worker1, Worker2, func(demoData *DemoData) {
		logrus.Info(demoData)
		count += 1
	})

	if count < 10 {
		t.Fail()
	}
}