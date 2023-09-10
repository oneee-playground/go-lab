package pipeline_test

import (
	"testing"

	to_stream "github.com/onee-only/go-lab/pattern/concurrency/channel/common/tostream"
	p "github.com/onee-only/go-lab/pattern/concurrency/channel/pipeline"
	"github.com/stretchr/testify/assert"
)

func TestPipe(t *testing.T) {
	add2 := func(val int) int { return val + 2 }
	values := []int{3, 3, 3, 4, 5}

	done := make(chan struct{})
	defer close(done)

	stream := p.Pipe(add2, done,
		p.Pipe(add2, done,
			p.Pipe(add2, done, to_stream.ToStream(done, values...)),
		),
	)

	cnt := 0
	for value := range stream {
		assert.Equal(t, values[cnt]+6, value)
		cnt++
	}
}
