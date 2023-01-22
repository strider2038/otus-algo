package structs

import (
	"fmt"
	"math/bits"
	"strings"
)

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

func (b BitSet) And(with BitSet) BitSet {
	bits1 := b
	bits2 := with

	if len(bits1) > len(bits2) {
		bits1, bits2 = bits2, bits1
	}

	and := make(BitSet, len(bits2))
	for i := 0; i < len(bits1); i++ {
		and[i] = bits1[i] & bits2[i]
	}

	return and
}

func (b BitSet) Or(with BitSet) BitSet {
	bits1 := b
	bits2 := with

	if len(bits1) > len(bits2) {
		bits1, bits2 = bits2, bits1
	}

	or := make(BitSet, len(bits2))
	for i := 0; i < len(bits1); i++ {
		or[i] = bits1[i] | bits2[i]
	}
	for i := len(bits1); i < len(bits2); i++ {
		or[i] = bits2[i]
	}

	return or
}

func (b BitSet) Intersects(with BitSet) bool {
	bits1 := b
	bits2 := with

	if len(bits1) > len(bits2) {
		bits1, bits2 = bits2, bits1
	}

	for i := 0; i < len(bits1); i++ {
		if bits1[i]&bits2[i] > 0 {
			return true
		}
	}

	return false
}

func (b BitSet) String() string {
	s := strings.Builder{}

	for i := len(b) - 1; i >= 0; i-- {
		s.WriteString(fmt.Sprintf("%x", b[i]))
	}

	return s.String()
}

func (b BitSet) split(n int) (int, int) {
	return n >> 6, n & 0x3F
}
