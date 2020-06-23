package pick

import (
	"container/heap"
	"sort"
)

type Ordered interface {
	Len() int
	Less(i, j int) bool
}

type Heap struct {
	indices []int
	data    Ordered
}

func (h *Heap) Len() int {
	return len(h.indices)
}

func (h *Heap) Swap(i, j int) {
	h.indices[i], h.indices[j] = h.indices[j], h.indices[i]
}

func (h *Heap) Less(i, j int) bool {
	return h.data.Less(h.indices[j], h.indices[i])
}

func (h *Heap) Push(x interface{}) {
	h.indices = append(h.indices, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := h.indices
	n := len(old)
	x := old[n-1]
	h.indices = old[0 : n-1]
	return x
}

func FirstN(data Ordered, n int) []int {
	h := &Heap{data: data}
	heap.Init(h)

	for i := 0; i < data.Len(); i++ {
		if i >= n {
			if h.data.Less(i, h.indices[0]) {
				heap.Pop(h)
				heap.Push(h, i)
			}
		} else {
			heap.Push(h, i)
		}
	}

	return h.indices
}
