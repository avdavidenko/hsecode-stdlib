package MaxQueue

import "container/list"
import "errors"

type MaxQueue struct {
	elements list.List
	maxes    list.List
}

func New() *MaxQueue {
	return &MaxQueue{}
}

func (q *MaxQueue) Max() (int, error) {
	if q.elements.Len() == 0 {
		return 0, errors.New("Queue is empty")
	}
	return q.maxes.Front().Value.(int), nil
}

func (q *MaxQueue) Pop() (int, error) {
	if q.elements.Len() == 0 {
		return 0, errors.New("Queue is empty")
	}

	temp := q.elements.Front()
	value := temp.Value.(int)
	q.elements.Remove(temp)

	if value == q.maxes.Front().Value {
		q.maxes.Remove(q.maxes.Front())
	}
	return value, nil
}

func (q *MaxQueue) Push(value int) {
	q.elements.PushBack(value)
	if q.maxes.Len() == 0 {
		q.maxes.PushBack(value)
	} else {
		for q.maxes.Len() > 0 && q.maxes.Back().Value.(int) < value {
			q.maxes.Remove(q.maxes.Back())
		}
		q.maxes.PushBack(value)
	}
}
