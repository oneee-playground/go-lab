package stack_test

import (
	"testing"

	"github.com/onee-only/go-lab/data-structure/linear/stack"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := stack.NewStack[int]()

	s.Push(1)
	s.Push(2)

	v, err := s.Pop()

	assert.NoError(t, err)
	assert.Equal(t, 2, v)
}
