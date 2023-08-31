package binaryinsertsort

import (
	"cmp"
	"slices"
)

func BinaryInsertSort[T cmp.Ordered](sorted []T, val T) []T {
	idx := locateIdx(sorted, val)

	return slices.Insert(sorted, idx, val)
}

func locateIdx[T cmp.Ordered](sorted []T, val T) int {
	if len(sorted) == 0 {
		return 0
	}

	mid := len(sorted) / 2
	if sorted[mid] < val {
		return locateIdx(sorted[mid+1:], val) + mid + 1
	} else {
		return locateIdx(sorted[:mid], val)
	}
}
