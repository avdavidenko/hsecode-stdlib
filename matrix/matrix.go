package matrix

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=int/dont_edit.go gen "ValueType=int"

type Matrix struct {
	Rows   int // number of rows
	Cols   int // number of columns
	matrix []ValueType
}

func New(n, m int) *Matrix {
	return &Matrix{
		Rows:   n,
		Cols:   m,
		matrix: make([]ValueType, m*n),
	}
}

func (M *Matrix) Get(i, j int) ValueType {
	if i > M.Rows || j > M.Cols {
		panic("error")
	}
	return M.matrix[i*M.Cols+j]
}

func (M *Matrix) Set(i, j int, v ValueType) {
	if i >= M.Rows || j >= M.Cols {
		panic("error")
	}
	M.matrix[i*M.Cols+j] = v
}
