package to_stream

func ToStream[T any](done chan struct{}, values ...T) chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)
		for _, value := range values {
			select {
			case <-done:
				return
			case stream <- value:
			}
		}
	}()

	return stream
}
