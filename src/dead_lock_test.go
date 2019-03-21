package src

import "testing"

func TestDeadLock(t *testing.T){
	o := make(chan struct{})

	o <- struct{}{}
}
