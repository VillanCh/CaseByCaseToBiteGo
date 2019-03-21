package demo

import (
	"math"
	"sync"
)

type SizedWaitGroup struct {
	Size    int
	current chan struct{}
	wg      sync.WaitGroup
}

func NewSizedWaitGroup(limit int) SizedWaitGroup {
	size := math.MaxInt32 // 2^32 - 1
	if limit > 0 {
		size = limit
	}
	return SizedWaitGroup{
		Size: size,

		current: make(chan struct{}, size),
		wg:      sync.WaitGroup{},
	}
}

func (s *SizedWaitGroup) Add() {
	select {
	case s.current <- struct{}{}:
		break
	}
	s.wg.Add(1)
}

func (s *SizedWaitGroup) Done() {
	<-s.current
	s.wg.Done()
}

func (s *SizedWaitGroup) Wait() {
	s.wg.Wait()
}

func (s *SizedWaitGroup) CountOfWaitingFor() int {
	return len(s.current)
}