package src

import (
	"github.com/sirupsen/logrus"
	"sync"
	"testing"
	"time"
)

func genIns() chan int {
	outC := make(chan int)
	go func() {
		defer close(outC)

		for i := 0; i <= 10; i ++ {
			outC <- i
		}
	}()

	return outC
}

func TestOneProducerMultiConsumers(t *testing.T){
	wg := sync.WaitGroup{}
	for result := range genIns(){
		wg.Add(1)

		go func(r int) {
			defer wg.Done()

			logrus.Infof("consumer fetch: %d", r)
			time.Sleep(500 * time.Millisecond)
		}(result)
	}

	wg.Wait()


	for result := range genIns(){
		// 总共有十个订阅者，订阅这一个结果
		for j := 0; j < 10; j ++ {
			wg.Add(1)

			go func(r int) {
				defer wg.Done()

				logrus.Infof("consumer fetch: %d", r)
				time.Sleep(500 * time.Millisecond)
			}(result)
		}
	}

	wg.Wait()
}