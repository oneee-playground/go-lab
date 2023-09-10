package ordone_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/onee-only/go-lab/pattern/concurrency/channel/ordone"
)

func createStream(done <-chan struct{}) <-chan struct{} {

	stream := make(chan struct{})
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case <-time.After(time.Millisecond * time.Duration(rand.Intn(1000))):
				stream <- struct{}{}
			}
		}
	}()

	return stream
}

func TestOrDone(t *testing.T) {

	done := make(chan struct{})

	src := createStream(done)

	go func() {
		<-time.After(time.Second * 2)
		close(done)
	}()

	results := ordone.OrDone(done, src)
	for range results {
	}
}
