package demo

import (
	"github.com/sirupsen/logrus"
	"time"
)

type DemoData struct {
	Type string
	Data interface{}
}

func Worker1(inC chan *DemoData, maxConcurrent int) chan *DemoData {

	nxtC := make(chan *DemoData)
	go func() {
		swg := NewSizedWaitGroup(maxConcurrent)
		defer close(nxtC)

		for data := range inC {
			swg.Add()
			go func(d *DemoData) {
				defer swg.Done()
				logrus.Infof("worker 1 handle data: %s", d)
				time.Sleep(1 * time.Second)
				nxtC <- &DemoData{
					Type: "worker1-result",
					Data: "worker1",
				}
			}(data)
		}

		swg.Wait()
	}()

	return nxtC
}

func Worker2(inC chan *DemoData, maxConcurrent int) chan *DemoData {

	nxtC := make(chan *DemoData)
	go func() {
		defer close(nxtC)

		swg := NewSizedWaitGroup(maxConcurrent)
		for data := range inC {
			swg.Add()
			go func(d *DemoData) {
				defer swg.Done()
				logrus.Infof("worker 2 handle data: %s", d)
				time.Sleep(1 * time.Second)
				nxtC <- &DemoData{
					Type: "worker2-result",
					Data: "worker2",
				}
			}(data)
		}

		swg.Wait()
	}()

	return nxtC
}

func Pipeline1(data chan *DemoData, worker1, worker2 func(chan *DemoData, int) chan *DemoData, fallback func(demoData *DemoData)) {

	// 连接管道
	p1 := worker1(data, 5)
	fallbackC := worker2(p1, 5)

	// fallback
	for result2 := range fallbackC {
		fallback(result2)
	}
}
