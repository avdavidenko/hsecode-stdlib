package bitset

import "math/bits"
import "errors"

type Bitset struct {
	length      int
	set         []uint64
	notUsedMask uint64
}

func New(size int) *Bitset {
	dataLength := size >> 6
	remainder := size & 63
	notUsedMask := uint64(0)
	if remainder != 0 {
		dataLength++
		notUsedMask = (^uint64(0)) << remainder
	}

	return &Bitset{length: size, set: make([]uint64, dataLength), notUsedMask: notUsedMask}
}

func (b *Bitset) All() bool {
	for i := 0; i < len(b.set)-1; i++ {
		if b.set[i] != (^uint64(0)) {
			return false
		}
	}

	if b.set[len(b.set)-1] != ^b.notUsedMask {
		return false
	}

	return true
}

func (b *Bitset) Any() bool {
	for i := 0; i < len(b.set); i++ {
		if b.set[i] != 0 {
			return true
		}
	}
	return false
}

func (b *Bitset) Count() int {
	count := 0
	for i := 0; i < len(b.set); i++ {
		count += bits.OnesCount64(b.set[i])
	}
	return count
}

func (b *Bitset) Flip() {
	for i := 0; i < len(b.set); i++ {
		b.set[i] = ^b.set[i]
	}
	b.set[len(b.set)-1] &= ^b.notUsedMask
}

func (b *Bitset) Reset() {
	for i := 0; i < len(b.set); i++ {
		b.set[i] = 0
	}
}

func (b *Bitset) Set(pos int, value bool) error {
	if pos < 0 || pos >= b.length {
		return errors.New("Out of range")
	}

	if value == true {
		b.set[pos>>6] |= 1 << (pos & 63)
	} else {
		b.set[pos>>6] &= ^(1 << (pos & 63))
	}

	return nil
}

func (b *Bitset) Test(pos int) (bool, error) {
	if pos < 0 || pos >= b.length {
		return false, errors.New("Out of range")
	}

	return (b.set[pos>>6]&(1<<(pos&63)) != 0), nil
}
