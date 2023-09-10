package fanin

import "sync"

func FanIn[T any](done <-chan struct{}, chans ...<-chan T) <-chan T {
	var wg sync.WaitGroup

	stream := make(chan T)

	for _, c := range chans {
		go func(ch <-chan T) {
			defer wg.Done()
			for val := range ch {
				select {
				case <-done:
					return
				case stream <- val:
				}
			}
		}(c)
	}

	wg.Add(len(chans))

	go func() {
		defer close(stream)
		wg.Wait()
	}()

	return stream
}
