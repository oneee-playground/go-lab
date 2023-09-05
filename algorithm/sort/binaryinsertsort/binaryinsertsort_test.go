package binaryinsertsort_test

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/onee-only/go-lab/algorithm/sort/binaryinsertsort"
	"github.com/stretchr/testify/assert"
)

func TestBinaryInsertSort(t *testing.T) {
	list := []int{}

	for i := 0; i < 50; i++ {
		list = binaryinsertsort.BinaryInsertSort(list, rand.Intn(50))
	}

	assert.True(t, slices.IsSorted(list))

}
