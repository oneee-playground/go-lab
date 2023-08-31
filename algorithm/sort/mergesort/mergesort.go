package mergesort

import "cmp"

func MergeSort[T cmp.Ordered](list []T) []T {
	if len(list) <= 1 {
		return list
	}

	mid := len(list) / 2
	left := MergeSort(list[:mid])
	right := MergeSort(list[mid:])

	return merge(left, right)
}

func merge[T cmp.Ordered](left []T, right []T) []T {

	var (
		leftIdx  int
		rightIdx int
		dstIdx   int

		dst = make([]T, len(left)+len(right))
	)

	insertLeft := func() {
		dst[dstIdx] = left[leftIdx]
		leftIdx++
	}
	insertRight := func() {
		dst[dstIdx] = right[rightIdx]
		rightIdx++
	}

	for leftIdx < len(left) || rightIdx < len(right) {
		switch {
		case leftIdx == len(left):
			insertRight()
		case rightIdx == len(right):
			insertLeft()
		default:
			if left[leftIdx] <= right[rightIdx] {
				insertLeft()
			} else {
				insertRight()
			}
		}
		dstIdx++
	}
	return dst
}
