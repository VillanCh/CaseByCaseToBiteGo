package src

import "testing"

func TestNoDeadLock(t *testing.T){

	o := make(chan struct{}, 1)

	o <- struct{}{}
}
