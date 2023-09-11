package bridge_test

import (
	"testing"

	"github.com/onee-only/go-lab/pattern/concurrency/channel/bridge"
)

func genStream() chan int {
	stream := make(chan int)

	go func() {
		defer close(stream)
		for i := 0; i < 10; i++ {
			stream <- i
		}
	}()

	return stream
}

func genSuperStream() <-chan <-chan int {
	superStream := make(chan (<-chan int))

	go func() {
		defer close(superStream)
		for i := 0; i < 4; i++ {
			superStream <- genStream()
		}
	}()

	return superStream
}

func TestBridge(t *testing.T) {

	done := make(chan struct{})
	superStream := genSuperStream()

	for range bridge.Bridge(done, superStream) {
	}

}
