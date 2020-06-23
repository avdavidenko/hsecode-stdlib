package pick

import (
	"sort"
)

type IntSlice []int

func (s IntSlice) Len() int {
	return len(s)
}

func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
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
	if left == right {
		return
	}

	pivotIndex := (left + right) / 2
	pivotIndex = partitionForNth(data, left, right, pivotIndex)

	if nth == pivotIndex {
		return
	}

	if nth < pivotIndex {
		selectForNth(data, left, pivotIndex-1, nth)
	} else {
		selectForNth(data, pivotIndex+1, right, nth)
	}
}

func NthElement(data sort.Interface, nth int) {
	if nth >= data.Len() || nth < 0 {
		panic("Out of range")
	}
	selectForNth(data, 0, data.Len()-1, nth)
}
