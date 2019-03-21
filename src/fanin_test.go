package src

import (
	"github.com/sirupsen/logrus"
	"sync"
	"testing"
)

func MultiProd() chan int {
	// 简单理解为多个生产者，一个消费者
	dataC := make(chan int)

	go func() {
		defer close(dataC)

		// 分配给多个生产者
		wg := sync.WaitGroup{}
		for i := 0; i < 100; i ++ {
			wg.Add(1)
			go func(data int) {
				defer wg.Done()

				dataC <- data
			}(i)
		}
		wg.Wait()

	}()

	return dataC
}

func TestFanIn(t *testing.T) {
	max := 0
	for result := range MultiProd(){
		if result > max{
			max = result
		}
		logrus.Infof("result: %d", result)
	}

	if max != 99 {
		t.Fail()
	}
}
