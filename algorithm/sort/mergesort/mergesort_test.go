package mergesort_test

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/onee-only/go-data-structures/algorithm/sort/mergesort"
	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	list := []int{}

	for i := 0; i < 50; i++ {
		list = append(list, rand.Intn(50))
	}

	assert.False(t, slices.IsSorted(list))

	list = mergesort.MergeSort(list)

	assert.True(t, slices.IsSorted(list))
}
