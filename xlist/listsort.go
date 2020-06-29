package xlist

import "container/list"

func getMiddleElement(begin, end *list.Element) *list.Element {
	slow := begin
	fast := begin

	for fast != end {
		fast = fast.Next()
		if fast != end {
			fast = fast.Next()
			slow = slow.Next()
		}
	}
	return slow
}

func Sort(data *list.List, less func(a, b *list.Element) bool) {
	if data.Len() < 2 {
		return
	}
	mergeSort(data, less, data.Front(), nil)
}

func mergeSort(
	data *list.List,
	less func(a, b *list.Element) bool,
	begin, end *list.Element) *list.Element {

	if begin == end || begin.Next() == end {
		return begin
	}

	middle := getMiddleElement(begin, end)
	begin = mergeSort(data, less, begin, middle)
	middle = mergeSort(data, less, middle, end)

	pos1 := begin
	pos2 := middle
	head := begin
	for pos2 != end {
		if less(pos1, pos2) {
			pos1 = pos1.Next()
		} else {
			tmp := pos2
			pos2 = pos2.Next()
			data.MoveBefore(tmp, pos1)
			if head == pos1 {
				head = tmp
			}
		}
	}
	return head
}
