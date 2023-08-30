package sparseset_test

import (
	"testing"

	"github.com/onee-only/go-data-structures/data-structure/common/kv"
	"github.com/onee-only/go-data-structures/data-structure/non-linear/sparseset"
	"github.com/stretchr/testify/assert"
)

func TestSparseSetAdd(t *testing.T) {
	s := sparseset.NewSparseSet[string, int]()

	dataset := []kv.KV[string, int]{
		{Key: "aaa", Value: 50},
		{Key: "bbb", Value: 20},
		{Key: "ccc", Value: 120},
	}

	for _, data := range dataset {
		s.Add(data.Key, data.Value)
	}

	for _, data := range dataset {
		v, ok := s.Get(data.Key)
		assert.True(t, ok)
		assert.Equal(t, data.Value, v)
	}

	v, ok := s.Get("wrong one")
	assert.False(t, ok)
	assert.Zero(t, v)
}

func TestSparseSetRemove(t *testing.T) {
	s := sparseset.NewSparseSet[string, int]()

	dataset := []kv.KV[string, int]{
		{Key: "aaa", Value: 50},
		{Key: "bbb", Value: 20},
		{Key: "ccc", Value: 120},
	}

	for _, data := range dataset {
		s.Add(data.Key, data.Value)
	}

	ok := s.Remove("bbb")
	assert.True(t, ok)

	v, ok := s.Get("bbb")
	assert.False(t, ok)
	assert.Zero(t, v)

	ok = s.Remove("wrong one")
	assert.False(t, ok)
}

func TestSparseSetIterator(t *testing.T) {
	s := sparseset.NewSparseSet[string, int]()

	dataset := []kv.KV[string, int]{
		{Key: "aaa", Value: 50},
		{Key: "bbb", Value: 20},
		{Key: "ccc", Value: 120},
	}

	for _, data := range dataset {
		s.Add(data.Key, data.Value)
	}

	in := func(val kv.KV[string, int], list []kv.KV[string, int]) bool {
		for _, elem := range list {
			if elem == val {
				return true
			}
		}
		return false
	}

	for i := s.Iterator(); !i.IsEnd(); i.Next() {
		assert.True(t, in(i.Get(), dataset))
	}
}
