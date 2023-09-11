package tee_test

import (
	"sync"
	"testing"

	to_stream "github.com/onee-only/go-lab/pattern/concurrency/channel/common/tostream"
	"github.com/onee-only/go-lab/pattern/concurrency/channel/tee"
)

func TestTee(t *testing.T) {
	done := make(chan struct{})

	stream := to_stream.ToStream(done, []int{1, 3, 2, 4, 2, 5}...)

	out1, out2 := tee.Tee(done, stream)
	var wg sync.WaitGroup

	read := func(ch <-chan int) {
		for range ch {
		}
		wg.Done()
	}

	wg.Add(2)

	go read(out1)
	go read(out2)

	wg.Wait()

}
