package pipeline_test

import (
	"testing"

	p "github.com/onee-only/go-lab/pattern/concurrency/channel/pipeline"
	"github.com/stretchr/testify/assert"
)

func toStream[T any](done chan struct{}, values ...T) chan T {
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

func TestPipe(t *testing.T) {
	add2 := func(val int) int { return val + 2 }
	values := []int{3, 3, 3, 4, 5}

	done := make(chan struct{})
	defer close(done)

	stream := p.Pipe(add2, done,
		p.Pipe(add2, done,
			p.Pipe(add2, done, toStream(done, values...)),
		),
	)

	cnt := 0
	for value := range stream {
		assert.Equal(t, values[cnt]+6, value)
		cnt++
	}
}
