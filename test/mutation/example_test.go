package mutation_test

import (
	"testing"

	"github.com/onee-only/go-lab/test/mutation"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	var (
		a = 50
		b = 0
	)

	sum := mutation.Sum(a, b)

	assert.Equal(t, a+b, sum)
}
