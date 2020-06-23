package vector

type Vector struct {
	Len    int
	vector []int
}

func New(cap int) *Vector {
	newCap := cap
	if newCap < 0 {
		newCap = 0
	}
	return &Vector{vector: make([]int, newCap), Len: 0}
}

func (a *Vector) Push(x int) {
	a.Insert(a.Len, x)
}

func (a *Vector) Pop() int {
	m := a.Get(a.Len - 1)
	a.Delete(a.Len - 1)
	return m
}

func (a *Vector) Set(idx int, x int) {
	if idx < 0 || idx >= a.Len {
		panic("Index out of range")
	}
	a.vector[idx] = x
}

func (a *Vector) Get(idx int) int {
	if idx < 0 || idx >= a.Len {
		panic("Index out of range")
	}
	return a.vector[idx]
}

func (a *Vector) Insert(idx int, x int) {
	if idx < 0 || idx > a.Len {
		panic("Index out of range")
	}

	if a.Len == len(a.vector) {
		newCap := a.Len * 2
		if newCap <= 0 {
			newCap = 8
		}
		newVector := make([]int, newCap)
		copy(newVector, a.vector)
		a.vector = newVector
	}

	if idx == a.Len {
		a.vector[idx] = x
	} else {
		for i := a.Len; i > idx; i-- {
			a.vector[i] = a.vector[i-1]
		}
		a.vector[idx] = x
	}

	a.Len++
}

func (a *Vector) Delete(idx int) {
	if idx < 0 || idx >= a.Len {
		panic("Index out of range")
	}

	if a.Len == 0 {
		panic("Vector is empty")
	}

	for i := idx + 1; i < a.Len; i++ {
		a.vector[i-1] = a.vector[i]
	}
	a.Len--

	if a.Len <= len(a.vector)/4 {
		newVector := make([]int, len(a.vector)/2)
		copy(newVector, a.vector)
		a.vector = newVector
	}
}
