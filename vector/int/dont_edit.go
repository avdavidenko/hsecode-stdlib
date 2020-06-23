package vector

type Vector struct {
	vector []int
}

func New(cap int) *Vector {
	newCap := cap
	if newCap < 0 {
		newCap = 0
	}
	return &Vector{vector: make([]int, 0, newCap)}
}

func (a *Vector) Push(x int) {
	a.Insert(len(a.vector), x)
}

func (a *Vector) Pop() int {
	m := a.Get(len(a.vector) - 1)
	a.Delete(len(a.vector) - 1)
	return m
}

func (a *Vector) Set(idx int, x int) {
	if idx < 0 || idx >= len(a.vector) {
		panic("Index out of range")
	}
	a.vector[idx] = x
}

func (a *Vector) Get(idx int) int {
	if idx < 0 || idx >= len(a.vector) {
		panic("Index out of range")
	}
	return a.vector[idx]
}

func (a *Vector) Insert(idx int, x int) {
	if idx < 0 || idx > len(a.vector) {
		panic("Index out of range")
	}

	if len(a.vector) == cap(a.vector) {
		newCap := len(a.vector) * 2
		if newCap <= 0 {
			newCap = 8
		}
		newVector := make([]int, len(a.vector), newCap)
		copy(newVector, a.vector)
		a.vector = newVector
	}

	if idx == len(a.vector) {
		a.vector = append(a.vector, x)
	} else {
		a.vector = append(a.vector, a.vector[len(a.vector)-1])
		for i := len(a.vector) - 1; i > idx; i-- {
			a.vector[i] = a.vector[i-1]
		}
		a.vector[idx] = x
	}
}

func (a *Vector) Delete(idx int) {
	if idx < 0 || idx >= len(a.vector) {
		panic("Index out of range")
	}

	if len(a.vector) == 0 {
		panic("Vector is empty")
	}

	for i := idx + 1; i < len(a.vector); i++ {
		a.vector[i-1] = a.vector[i]
	}
	a.vector = a.vector[:len(a.vector)-1]

	if len(a.vector) <= cap(a.vector)/4 {
		newVector := make([]int, len(a.vector), cap(a.vector)/2)
		copy(newVector, a.vector)
		a.vector = newVector
	}
}
