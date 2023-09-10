package pipeline

func Pipe[T any](f func(val T) T, done <-chan struct{}, values <-chan T) chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)

		for val := range values {
			select {
			case <-done:
				return
			case stream <- f(val):
			}
		}
	}()

	return stream
}
