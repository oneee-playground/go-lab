package skiplist_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/onee-only/go-lab/data-structure/linear/list/skiplist"
	"github.com/stretchr/testify/assert"
)

func TestSkipList(t *testing.T) {
	sl := skiplist.NewSkipList[int, int](50)

	for i := 0; i < 500; i++ {
		sl.Set(i, i)
	}

	val, err := sl.Find(250)
	assert.NoError(t, err)
	assert.Equal(t, 250, val)
}

func TestSkipListRemove(t *testing.T) {
	sl := skiplist.NewSkipList[int, int](5)

	for i := 0; i < 50; i++ {
		sl.Set(i, i)
	}

	sl.Remove(32)

	_, err := sl.Find(32)
	assert.Error(t, err)
}

func BenchmarkSkipList(b *testing.B) {

	go func() {
		<-time.After(time.Minute * 2)
		b.Fail()
	}()

	sl := skiplist.NewSkipList[int, int](20)

	for i := 0; i < b.N; i++ {
		sl.Set(i, rand.Int())
	}
}
