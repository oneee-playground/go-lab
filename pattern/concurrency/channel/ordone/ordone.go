package ordone

func OrDone[T any](done <-chan struct{}, ch <-chan T) <-chan T {

	stream := make(chan T)
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-ch:
				if !ok {
					return
				}
				select {
				case stream <- v:
				case <-done:
				}
			}
		}
	}()

	return stream
}
