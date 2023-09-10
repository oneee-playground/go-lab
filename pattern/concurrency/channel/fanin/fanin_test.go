package fanin_test

import (
	"math/rand"
	"testing"

	to_stream "github.com/onee-only/go-lab/pattern/concurrency/channel/common/tostream"
	"github.com/onee-only/go-lab/pattern/concurrency/channel/fanin"
	"github.com/stretchr/testify/assert"
)

func genRandNum(take int) (list []int) {
	for i := 0; i < take; i++ {
		list = append(list, rand.Int())
	}
	return list
}

func TestFanIn(t *testing.T) {

	done := make(chan struct{})
	defer close(done)

	stream1 := to_stream.ToStream(done, genRandNum(30)...)
	stream2 := to_stream.ToStream(done, genRandNum(20)...)
	stream3 := to_stream.ToStream(done, genRandNum(50)...)

	stream := fanin.FanIn(done, stream1, stream2, stream3)

	cnt := 0
	for range stream {
		cnt++
	}
	assert.Equal(t, 100, cnt)
}
