package tree

import "sort"

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func NewBST(elements []int) *Tree {
	h := make(map[int]struct{})

	unique := make([]int, 0)

	for _, e := range elements {
		_, ok := h[e]
		if !ok {
			h[e] = struct{}{}
			unique = append(unique, e)
		}
	}

	sort.Ints(unique)

	return fromSorted(unique)
}

func (T *Tree) IsSym() bool {
	return isMirror(T, T)
}

func (T *Tree) InOrder(visit func(node *Tree)) {
	stack := make([](*Tree), 0)
	curr := T

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visit(curr)

		curr = curr.Right
	}
}

func fromSorted(unique []int) *Tree {
	if len(unique) == 0 {
		return nil
	}

	if len(unique) == 1 {
		return &Tree{Value: unique[0], Left: nil, Right: nil}
	}

	med := len(unique) / 2
	return &Tree{Value: unique[med], Left: fromSorted(unique[:med]), Right: fromSorted(unique[med+1:])}
}

func isMirror(T1, T2 *Tree) bool {
	if T1 == nil && T2 == nil {
		return true
	}

	if T1 == nil || T2 == nil {
		return false
	}

	if T1.Value != T2.Value {
		return false
	}

	return isMirror(T1.Left, T2.Right) && isMirror(T1.Right, T2.Left)
}
