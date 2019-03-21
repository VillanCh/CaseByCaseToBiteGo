package src

import (
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestChanVsSharing(t *testing.T) {
	srcC := make(chan int)
	go func() {
		defer close(srcC)
		for i := 0; i < 5; i ++ {
			srcC <- i
		}
	}()


	// 这里的代码可能存在问题？
	//    怎么解决？
	for e := range srcC {
		go func() {
			logrus.Infof("goroutine fetch: %d", e)
		}()
	}

	for j := 0; j < 10; j ++ {
		go func() {
			logrus.Infof("goroutine fetch: %d", j)
		}()
	}

	time.Sleep(1 * time.Second)
}
