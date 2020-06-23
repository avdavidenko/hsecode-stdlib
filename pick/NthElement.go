package pick

import (
	"math/rand"
	"sort"
)

type IntSlice []int

func (s IntSlice) Len() int {
	return len(s)
}

func (s IntSlice) Swap(i, j int) {
	t := s[i]
	s[i] = s[j]
	s[j] = t
}

func (s IntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func partitionForNth(data sort.Interface, left int, right int, pivotIndex int) int {
	data.Swap(pivotIndex, right)
	storeIndex := left
	for i := left; i < right; i++ {
		if data.Less(i, right) {
			data.Swap(storeIndex, i)
			storeIndex++
		}
	}

	data.Swap(storeIndex, right)
	return storeIndex
}

func selectForNth(data sort.Interface, left int, right int, nth int) {
	for left < right {
		pivotIndex := rand.Intn(right+1-left) + left
		pivotIndex = partitionForNth(data, left, right, pivotIndex)

		if nth == pivotIndex {
			return
		}

		if nth < pivotIndex {
			right = pivotIndex - 1
		} else {
			left = pivotIndex + 1
		}
	}
}

func NthElement(data sort.Interface, nth int) {
	if nth >= data.Len() || nth < 0 {
		panic("Out of range")
	}
	selectForNth(data, 0, data.Len()-1, nth)
}
