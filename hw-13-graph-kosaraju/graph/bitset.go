package graph

import "math/bits"

type BitSet []uint64

func NewBitSet(maxN int) BitSet {
	return make(BitSet, (maxN>>6)+1)
}

func (b BitSet) Set(n int) {
	i, j := b.split(n)

	b[i] = b[i] | (1 << j)
}

func (b BitSet) Unset(n int) {
	i, j := b.split(n)

	b[i] = b[i] & ^(1 << j)
}

func (b BitSet) IsSet(n int) bool {
	i, j := b.split(n)

	return b[i]&(1<<j) != 0
}

func (b BitSet) OnesCount() int {
	count := 0
	for _, u := range b {
		count += bits.OnesCount64(u)
	}
	return count
}

func (b BitSet) split(n int) (int, int) {
	return n >> 6, n & 0x3F
}
