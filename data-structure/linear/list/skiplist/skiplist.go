package skiplist

import (
	"cmp"
	"errors"
	"math/rand"

	"github.com/onee-only/go-lab/data-structure/linear/stack"
)

type SkipList[TKey cmp.Ordered, TValue any] struct {
	head     *header[TKey, TValue]
	maxLevel int
}

func NewSkipList[TKey cmp.Ordered, TValue any](maxLevel int) *SkipList[TKey, TValue] {
	return &SkipList[TKey, TValue]{
		head:     newHeader[TKey, TValue](*new(TKey), *new(TValue), maxLevel),
		maxLevel: maxLevel,
	}
}

func (s *SkipList[TKey, TValue]) Find(key TKey) (TValue, error) {

	var (
		level     = s.maxLevel - 1
		next, cur = s.head, s.head
	)

	for {
		next = cur.next[level]
		if next == nil || next.data.Key < key {
			if level == 0 {
				return *new(TValue), errors.New("not found")
			}

			level--
			continue
		}

		if next.data.Key == key {
			return next.data.Value, nil
		}

		cur = next
	}
}

func (s *SkipList[TKey, TValue]) Set(key TKey, value TValue) {

	var (
		prevTops = stack.NewStack[*header[TKey, TValue]]()

		headrLevel = genRandomLvl(s.maxLevel)
		level      = s.maxLevel - 1
		next, cur  = s.head, s.head
	)

	headr := newHeader[TKey, TValue](key, value, headrLevel)
	if cur.next[0] == nil {
		for idx := range headr.next {
			cur.next[idx] = headr
		}

		return
	}

	for {
		next = cur.next[level]
		if next == nil || next.data.Key < key {
			prevTops.Push(cur)

			if level == 0 {
				for idx := range headr.next {
					if idx >= len(next.next) {
						cur, _ = prevTops.Pop()
					}

					headr.next[idx] = cur.next[idx]
					cur.next[idx] = headr
				}
				return
			}

			level--
			continue
		}

		if next.data.Key == key {
			next.data.Value = value
			return
		}

		cur = next
	}

}

func (s *SkipList[TKey, TValue]) Remove(key TKey) error {

	var (
		prevTops = stack.NewStack[*header[TKey, TValue]]()

		level     = s.maxLevel - 1
		next, cur = s.head, s.head
	)

	for {
		next = cur.next[level]
		if next == nil || next.data.Key < key {
			prevTops.Push(cur)

			if level == 0 {
				return errors.New("not found")
			}

			level--
			continue
		}

		if next.data.Key == key {
			for idx := range next.next {
				nextOne := next.next[idx]
				if idx < len(cur.next) {
					cur.next[idx] = nextOne
				} else {
					prevTop, _ := prevTops.Pop()

					prevTop.next[idx] = nextOne
				}
			}
		}

		cur = next
	}
}

func genRandomLvl(max int) int {
	for i := 1; i < max; i++ {
		if rand.Intn(2) == 1 {
			return i
		}
	}
	return max
}
