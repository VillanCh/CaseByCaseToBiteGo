package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime"
	"time"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	ticker := time.NewTicker(1 * time.Second)
	var m runtime.MemStats

	runtime.ReadMemStats(&m)
	logrus.Println("test normal")
	logrus.Printf("\tAlloc = %v MiB", bToMb(m.Alloc))
	logrus.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	logrus.Printf("\tSys = %v MiB", bToMb(m.Sys))
	logrus.Printf("\tNumGC = %v\n", m.NumGC)

	count := 1000
	for i := 0; i < count; i ++ {
		//logrus.Infof("start %d/%d", i+1, count)
		c := make(chan struct{})
		go func(testC chan struct{}) {
			defer close(c)

			for k := 0; k < 10000; k ++ {
				time.Sleep(100 * time.Millisecond)
			}

			c <- struct{}{}

			http.Get("http://127.0.0.1")
			time.Sleep(1000 * time.Second)
		}(c)
	}

	logrus.Println("test normal")

	for {
		select {
		case <-ticker.C:
			runtime.ReadMemStats(&m)
			logrus.Printf("Alloc = %v MiB", bToMb(m.Alloc))
			logrus.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
			logrus.Printf("\tSys = %v MiB", bToMb(m.Sys))
			logrus.Printf("\tNumGC = %v\n", m.NumGC)
		}
	}
}
