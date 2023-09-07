package skiplist

import (
	"cmp"

	"github.com/onee-only/go-lab/data-structure/common/kv"
)

type header[TKey cmp.Ordered, TValue any] struct {
	next []*header[TKey, TValue]
	data *kv.KV[TKey, TValue]
}

func newHeader[TKey cmp.Ordered, TValue any](key TKey, value TValue, level int) *header[TKey, TValue] {
	return &header[TKey, TValue]{
		data: &kv.KV[TKey, TValue]{
			Key:   key,
			Value: value,
		},
		next: make([]*header[TKey, TValue], level),
	}
}
