package demo

import (
	"github.com/sirupsen/logrus"
	"time"
)

// Worker 基础类和定义
type WorkerIf interface {
	Handle(data *DemoData, outChan chan *DemoData) error
	GetName() string
}

type WorkerBase struct {
	Name string
}

func (w *WorkerBase) GetName() string {
	return w.Name
}

// RunWorker 负责把一个 Worker 变成真的消费者兼生产者，并且可以控制并发
func RunWorker(inC chan *DemoData, worker WorkerIf, maxConcurrent int) chan *DemoData {
	nxtC := make(chan *DemoData)
	go func() {
		// 作为生产者，一定要关闭管道
		defer close(nxtC)

		// 限制消费者的消费能力
		swg := NewSizedWaitGroup(maxConcurrent)
		for data := range inC {

			// 超出可成熟的并发数时会阻塞
			swg.Add()
			go func(d *DemoData) {
				defer swg.Done()

				err := worker.Handle(d, nxtC)
				if err != nil {
					logrus.Errorf("worker: %s error: %s", worker.GetName(), err)
				}

			}(data)
		}

		// 这里 Wait 确保所有任务都已经完成了
		swg.Wait()
	}()

	return nxtC
}

// 定义 Pipeline 中的第一个中间单元，既是消费者也是生产者
type Worker1Pro struct {
	WorkerBase
}

func NewWorker1Pro() *Worker1Pro {
	// 在这里进行各种配置
	w := &Worker1Pro{}
	w.Name = "worker1-pro"
	return w
}

func (w1 *Worker1Pro) Handle(data *DemoData, outChan chan *DemoData) error {
	logrus.Infof("worker1 pro recv data: %s", data)
	time.Sleep(1 * time.Second)
	outChan <- &DemoData{
		Type: "worker1-pro-result",
		Data: "worker 1 pro result string....",
	}
	return nil
}

// 定义第二个中间单元，既是消费者也是生产者
type Worker2Pro struct {
	WorkerBase
}

func NewWorker2Pro() *Worker2Pro {
	// 在这里进行各种配置
	w := &Worker2Pro{}
	w.Name = "worker2-pro"
	return w
}

func (w1 *Worker2Pro) Handle(data *DemoData, outChan chan *DemoData) error {
	logrus.Infof("worker2 pro recv data: %s", data)
	time.Sleep(1 * time.Second)
	outChan <- &DemoData{
		Type: "worker2-pro-result",
		Data: "worker 2 pro result string....",
	}
	return nil
}

// 自动构建管道
func PipelinePro(originChan chan *DemoData, workers ...WorkerIf) chan *DemoData {
	lastC := originChan
	for _, worker := range workers {
		lastC = RunWorker(lastC, worker, 5)
	}
	return lastC
}
