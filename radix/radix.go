package radix

func ExtractByte(value uint64, n int) int {
	return int((value >> (n << 3)) & 0b11111111)
}

func Sort(data []uint64) {
	temp := make([]uint64, len(data))

	source := &temp
	target := &data
	for i := 0; i < 8; i++ {
		source, target = target, source
		bins := make([]int, 257)
		for j := 0; j < len(*source); j++ {
			bins[ExtractByte((*source)[j], i)+1]++
		}

		if bins[0] == len(*source) {
			source, target = target, source
		} else {

			fmt.Println(bins)
			for j := 1; j < 256; j++ {
				bins[j] += bins[j-1]
			}

			fmt.Println(bins)
			for j := 0; j < len(*source); j++ {
				byteVal := ExtractByte((*source)[j], i)
				(*target)[bins[byteVal]] = (*source)[j]
				bins[byteVal]++
			}
		}
	}

	if target != &data {
		copy(data, *target)
	}
}
