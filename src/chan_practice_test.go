package src

import (
	"testing"
)

func TestChanPractice(t *testing.T) {
	modC := make(chan struct{})
	go func() {
		// 保证生产者关闭
		defer close(modC)

		t.Log("生产者 输入一个空结构体")
		modC <- struct{}{}
	}()

	// use ur modC
	t.Log("消费者收到 ", <- modC)

	t.Fail()
}
