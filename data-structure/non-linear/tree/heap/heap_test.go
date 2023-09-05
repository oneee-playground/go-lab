package heap_test

import (
	"math/rand"
	"testing"

	"github.com/onee-only/go-lab/data-structure/non-linear/tree/heap"
	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	h := heap.NewHeap[int]()

	for i := 0; i < 20; i++ {
		h.Push(rand.Intn(50))
	}

	min, _ := h.Pop()
	for h.Size() != 0 {
		val, _ := h.Pop()
		assert.LessOrEqual(t, min, val)
	}
}

func BenchmarkHeapPush(b *testing.B) {

	h := heap.NewHeap[int]()
	for i := 0; i < b.N; i++ {
		h.Push(rand.Intn(50))
	}
}
