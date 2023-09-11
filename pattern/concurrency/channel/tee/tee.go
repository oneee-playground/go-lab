package tee

import "github.com/onee-only/go-lab/pattern/concurrency/channel/ordone"

func Tee[T any](done <-chan struct{}, in <-chan T) (_, _ <-chan T) {
	var (
		out1 = make(chan T)
		out2 = make(chan T)
	)

	go func() {
		defer close(out1)
		defer close(out2)

		for val := range ordone.OrDone(done, in) {
			var out1, out2 = out1, out2
			for i := 0; i < 2; i++ {
				select {
				case <-done:
				case out1 <- val:
					out1 = nil
				case out2 <- val:
					out2 = nil
				}
			}
		}
	}()

	return out1, out2
}
