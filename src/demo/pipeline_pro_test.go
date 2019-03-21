package demo

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestPipelinePro(t *testing.T) {
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

	w1 := NewWorker1Pro()
	w2 := NewWorker2Pro()

	count := 0
	for result := range PipelinePro(origin, w1, w2) {
		count += 1
		logrus.Infof("Pipeline Pro Fallback: %s", result)
	}

	t.Fail()


}
