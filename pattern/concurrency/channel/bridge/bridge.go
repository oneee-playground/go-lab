package bridge

import "github.com/onee-only/go-lab/pattern/concurrency/channel/ordone"

func Bridge[T any](done <-chan struct{}, chanStream <-chan <-chan T) <-chan T {

	outStream := make(chan T)

	go func() {
		defer close(outStream)

		for {
			var stream <-chan T
			select {
			case candidate, ok := <-chanStream:
				if !ok {
					return
				}
				stream = candidate
			case <-done:
				return
			}
			for val := range ordone.OrDone(done, stream) {
				select {
				case outStream <- val:
				case <-done:
				}
			}
		}
	}()

	return outStream
}
