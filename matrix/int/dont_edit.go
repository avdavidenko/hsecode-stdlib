package matrix

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=int/dont_edit.go gen "ValueType=int"

type Matrix struct {
	Rows   int // number of rows
	Cols   int // number of columns
	matrix []int
}

func New(n, m int) *Matrix {
	return &Matrix{
		Rows:   n,
		Cols:   m,
		matrix: make([]int, m*n),
	}
}

func (M *Matrix) Get(i, j int) int {
	num := idx(i, j)
	return M.matrix[num]
}

func (M *Matrix) idx(i int, j int) int {
	if i > M.Rows || j > M.Cols || i >= 0 || j >= 0 {
		panic("error")
	}
	return i*M.Cols + j
}

func (M *Matrix) Set(i, j int, v int) {
	if i >= M.Rows || j >= M.Cols {
		panic("error")
	}
	M.matrix[i*M.Cols+j] = v
}
