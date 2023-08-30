package heap

import (
	"cmp"
	"errors"
)

// Heap is a binary tree whose node is always greater or less than its parent node.
//
// This heap implements min heap tree where its root node has the minimum value.
type Heap[T cmp.Ordered] struct {
	nodes []T
}

// NewHeap creates new heap
func NewHeap[T cmp.Ordered]() *Heap[T] {
	return &Heap[T]{
		nodes: []T{},
	}
}

// Size returns size of heap's nodes
func (h *Heap[T]) Size() int { return len(h.nodes) }

// Clear clears nodes of heap
func (h *Heap[T]) Clear() { h.nodes = []T{} }

// Push pushes element to heap. then it calls upheapify to sort elements
func (h *Heap[T]) Push(value T) {
	h.nodes = append(h.nodes, value)

	h.upheapify(h.Size() - 1)
}

// Pop pops the element out of heap. then it calls downheapify to sort elements
func (h *Heap[T]) Pop() (value T, err error) {

	if h.Size() == 0 {
		return value, errors.New("the heap is empty")
	}

	value = h.nodes[0]
	h.swap(0, h.Size()-1)
	h.nodes = h.nodes[:h.Size()-1]

	h.downheapify(0)

	return value, nil
}

func (h *Heap[T]) upheapify(idx int) {

	if idx == 0 {
		return
	}

	if parentIdx := parent(idx); h.nodes[idx] < h.nodes[parentIdx] {
		h.swap(parentIdx, idx)
		h.upheapify(parentIdx)
	}
}

func (h *Heap[T]) downheapify(idx int) {

	children := []int{leftChild(idx), rightChild(idx)}

	for _, childIdx := range children {
		if childIdx < h.Size() && h.nodes[idx] > h.nodes[childIdx] {
			h.swap(idx, childIdx)
			h.downheapify(childIdx)
			break
		}
	}
}

func (h *Heap[T]) swap(i, j int) { h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i] }

func parent(idx int) int { return (idx - 1) / 2 }

func leftChild(idx int) int { return (idx * 2) + 1 }

func rightChild(idx int) int { return (idx * 2) + 2 }
