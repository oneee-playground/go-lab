package sort_test

import (
	"math/rand"
	"testing"

	"github.com/onee-only/go-data-structures/algorithm/sort/mergesort"
	"github.com/onee-only/go-data-structures/algorithm/sort/quicksort"
)

func unorderedList(n int) (list []int) {
	for i := 0; i < n; i++ {
		list = append(list, rand.Int())
	}
	return
}

func orderedList(n int) (list []int) {
	for i := 0; i < n; i++ {
		list = append(list, i)
	}
	return
}

func BenchmarkQuickSort(b *testing.B) {
	list := unorderedList(b.N)
	b.ResetTimer()

	quicksort.QuickSort(list)
}

func BenchmarkQuickSortWorst(b *testing.B) {
	list := orderedList(b.N)
	b.ResetTimer()

	quicksort.QuickSort(list)
}

func BenchmarkMergeSort(b *testing.B) {
	list := unorderedList(b.N)
	b.ResetTimer()

	mergesort.MergeSort(list)
}
