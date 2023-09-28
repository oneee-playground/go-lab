package nestedzerovalue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type parent struct {
	child child
}

type child struct {
	a int
}

func TestNestedZeroValue(t *testing.T) {
	var p parent

	assert.Zero(t, p)
	assert.Zero(t, p.child)
	assert.Zero(t, p.child.a)

}
