package quicksort

import (
	"cmp"
)

func QuickSort[T cmp.Ordered](list []T) {
	if len(list) <= 1 {
		return
	}

	idx := partition(list)

	QuickSort(list[:idx])
	QuickSort(list[idx+1:])
}

func partition[T cmp.Ordered](list []T) int {

	var (
		pivot    T
		leftIdx  int = 1
		rightIdx int = len(list) - 1
	)

	pivot = list[0]

	for {
		for leftIdx < len(list) && list[leftIdx] <= pivot {
			leftIdx++
		}

		for rightIdx > 0 && list[rightIdx] > pivot {
			rightIdx--
		}

		if leftIdx >= rightIdx {
			swap(list, 0, leftIdx-1)
			return leftIdx - 1
		}
		swap(list, leftIdx, rightIdx)
	}
}

func swap[T cmp.Ordered](list []T, i, j int) {
	list[j], list[i] = list[i], list[j]
}
