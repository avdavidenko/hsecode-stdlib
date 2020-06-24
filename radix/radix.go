package radix

func Sort(data []uint64) {
	temp := make([]uint64, len(data))

	source := &temp
	target := &data
	bins := make([]int, 257)
	for i := 0; i < 8; i++ {
		shift := i << 3
		source, target = target, source
		for j := 0; j < 257; j++ {
			bins[j] = 0
		}
		for _, value := range *source {
			byteVal := int((value >> shift) & 0b11111111)
			bins[byteVal+1]++
		}

		if bins[0] == len(*source) {
			source, target = target, source
		} else {

			for j := 1; j < 256; j++ {
				bins[j] += bins[j-1]
			}

			for _, value := range *source {
				byteVal := int((value >> shift) & 0b11111111)
				(*target)[bins[byteVal]] = value
				bins[byteVal]++
			}
		}
	}

	if target != &data {
		copy(data, *target)
	}
}
