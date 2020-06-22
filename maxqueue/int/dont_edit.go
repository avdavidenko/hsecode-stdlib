package MaxQueue

import "errors"

type MaxQueue struct {
	elements_in []int
	maxes_in    []int
	elements_ou []int
	maxes_ou    []int
}

func New() *MaxQueue {
	return &MaxQueue{}
}

func (q *MaxQueue) Max() (int, error) {
	if len(q.maxes_ou) > 0 && len(q.maxes_in) > 0 {
		if q.maxes_ou[len(q.maxes_ou)-1] > q.maxes_in[len(q.maxes_in)-1] {
			return q.maxes_ou[len(q.maxes_ou)-1], nil
		} else {
			return q.maxes_in[len(q.maxes_in)-1], nil
		}
	}

	if len(q.maxes_ou) > 0 {
		return q.maxes_ou[len(q.maxes_ou)-1], nil
	}

	if len(q.maxes_in) > 0 {
		return q.maxes_in[len(q.maxes_in)-1], nil
	}

	return 0, errors.New("Queue is empty")
}

func (q *MaxQueue) Pop() (int, error) {
	if len(q.maxes_ou) == 0 {
		for i := len(q.maxes_in) - 1; i >= 0; i-- {
			value := q.elements_in[i]
			q.elements_ou = append(q.elements_ou, value)
			max := value
			if len(q.maxes_ou) > 0 && q.maxes_ou[len(q.maxes_ou)-1] > value {
				max = q.maxes_ou[len(q.maxes_ou)-1]
			}
			q.maxes_ou = append(q.maxes_ou, max)
		}
		q.maxes_in = q.maxes_in[:0]
		q.elements_in = q.elements_in[:0]
	}

	if len(q.maxes_ou) == 0 {
		return 0, errors.New("Queue is empty")
	}

	q.maxes_ou = q.maxes_ou[:len(q.maxes_ou)-1]
	value := q.elements_ou[len(q.elements_ou)-1]
	q.elements_ou = q.elements_ou[:len(q.elements_ou)-1]

	return value, nil
}

func (q *MaxQueue) Push(value int) {
	q.elements_in = append(q.elements_in, value)
	if len(q.maxes_in) == 0 {
		q.maxes_in = append(q.maxes_in, value)
	} else {
		max := q.maxes_in[len(q.maxes_in)-1]
		if max > value {
			q.maxes_in = append(q.maxes_in, max)
		} else {
			q.maxes_in = append(q.maxes_in, value)
		}
	}
}
