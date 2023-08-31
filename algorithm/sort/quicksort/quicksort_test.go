package quicksort_test

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/onee-only/go-data-structures/algorithm/sort/quicksort"
	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	list := []int{}

	for i := 0; i < 50; i++ {
		list = append(list, rand.Intn(50))
	}

	assert.False(t, slices.IsSorted(list))

	quicksort.QuickSort(list)

	assert.True(t, slices.IsSorted(list))
}

func BenchmarkQuickSort(b *testing.B) {
	list := []int{}

	for i := 0; i < b.N; i++ {
		list = append(list, rand.Int())
	}

	b.ResetTimer()

	quicksort.QuickSort(list)
}

func BenchmarkQuickSortWorst(b *testing.B) {
	list := []int{}

	for i := 0; i < b.N; i++ {
		list = append(list, i)
	}

	b.ResetTimer()

	quicksort.QuickSort(list)
}
