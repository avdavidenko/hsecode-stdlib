package ndarray

type NDArray struct {
	shape []int
}

func New(shape ...int) *NDArray {
	if len(shape) == 0 {
		panic("invalid shape")
	}
	for i := 0; i < len(shape); i++ {
		if shape[i] < 0 {
			panic("invalid shape")
		}
	}
	return &NDArray{shape: shape}
}
func (nda *NDArray) Idx(indicies ...int) int {
	if len(indicies) != len(nda.shape) {
		panic("invalid length")
	}

	for i := 0; i < len(nda.shape); i++ {
		if indicies[i] < 0 || indicies[i] >= nda.shape[i] {
			panic("invalid index")
		}
	}

	idx := indicies[0]
	for i := 1; i < len(nda.shape); i++ {
		idx = idx*nda.shape[i] + indicies[i]
	}
	return idx
}
