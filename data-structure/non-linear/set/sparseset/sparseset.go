package sparseset

import (
	"cmp"

	"github.com/onee-only/go-lab/data-structure/common/iterator"
	"github.com/onee-only/go-lab/data-structure/common/kv"
)

type SparseSet[TKey cmp.Ordered, TValue any] struct {
	dense  []kv.KV[TKey, TValue]
	sparse map[TKey]int
}

func NewSparseSet[TKey cmp.Ordered, TValue any]() *SparseSet[TKey, TValue] {
	return &SparseSet[TKey, TValue]{
		sparse: make(map[TKey]int),
	}
}

func (s *SparseSet[TKey, TValue]) Add(key TKey, value TValue) {
	if idx, ok := s.sparse[key]; ok {
		s.dense[idx].Value = value
		return
	}

	s.dense = append(s.dense, kv.KV[TKey, TValue]{
		Key:   key,
		Value: value,
	})

	s.sparse[key] = len(s.dense) - 1
}

func (s *SparseSet[TKey, TValue]) Get(key TKey) (value TValue, found bool) {
	if idx, ok := s.sparse[key]; ok {
		value = s.dense[idx].Value
		found = true
	}
	return
}

func (s *SparseSet[TKey, TValue]) Remove(key TKey) (found bool) {
	if idx, ok := s.sparse[key]; ok {
		last := len(s.dense) - 1

		if idx < last {
			s.dense[idx] = s.dense[last]
			s.sparse[s.dense[last].Key] = idx
		}
		s.dense = s.dense[:last]
		delete(s.sparse, key)

		found = true
	}
	return
}

type SparseSetIterator[TKey cmp.Ordered, TValue any] struct {
	dense []kv.KV[TKey, TValue]
	idx   int
}

func (s *SparseSet[TKey, TValue]) Iterator() iterator.Iterator[kv.KV[TKey, TValue]] {
	return &SparseSetIterator[TKey, TValue]{
		dense: s.dense,
		idx:   0,
	}
}

func (i *SparseSetIterator[TKey, TValue]) Next() { i.idx++ }

func (i *SparseSetIterator[TKey, TValue]) IsEnd() bool { return i.idx >= len(i.dense) }

func (i *SparseSetIterator[TKey, TValue]) Get() kv.KV[TKey, TValue] { return i.dense[i.idx] }
