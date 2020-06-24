package bitset

import "math/bits"
import "errors"

type Bitset struct {
	length      int
	set         []uint8
	notUsedMask uint8
}

func New(size int) *Bitset {
	dataLength := size >> 3
	remainder := size & 0x7
	notUsedMask := uint8(0)
	if remainder != 0 {
		dataLength++
		notUsedMask = 0xFF << remainder
	}

	return &Bitset{length: size, set: make([]uint8, dataLength), notUsedMask: notUsedMask}
}

func (b *Bitset) All() bool {
	return b.Count() == b.length
}

func (b *Bitset) Any() bool {
	for _, value := range b.set {
		if value != 0 {
			return true
		}
	}
	return false
}

func (b *Bitset) Count() int {
	count := 0
	for _, value := range b.set {
		count += bits.OnesCount8(value)
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
		b.set[pos>>3] |= 1 << (pos & 7)
	} else {
		b.set[pos>>3] &= ^(1 << (pos & 7))
	}

	return nil
}

func (b *Bitset) Test(pos int) (bool, error) {
	if pos < 0 || pos >= b.length {
		return false, errors.New("Out of range")
	}

	return (b.set[pos>>3]&(1<<(pos&7)) != 0), nil
}
