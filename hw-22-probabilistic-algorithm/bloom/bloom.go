package bloom

import (
	"crypto/sha256"
	"encoding/binary"
	"math"

	"github.com/strider2038/otus-algo/pkg/structs"
)

type Filter struct {
	m    uint
	k    uint
	bits structs.BitSet
}

func NewFilter(n uint, p float64) *Filter {
	// M = -n * log2(p) / ln(2)
	m := uint(math.Ceil(-float64(n) * math.Log2(p) / math.Log(2)))
	// k = ln(1/p)
	k := uint(math.Ceil(math.Log(1 / p)))

	if m == 0 {
		m = 1
	}
	if k == 0 {
		k = 1
	}

	return &Filter{m: m, k: k, bits: structs.NewBitSet(int(m))}
}

func (f *Filter) Add(value string) {
	for i := uint(0); i < f.k; i++ {
		index := hash(value, i) % f.m
		f.bits.Set(int(index))
	}
}

func (f *Filter) Contains(value string) bool {
	for i := uint(0); i < f.k; i++ {
		index := hash(value, i) % f.m
		if !f.bits.IsSet(int(index)) {
			return false
		}
	}

	return true
}

func hash(value string, seed uint) uint {
	h := sha256.Sum256(binary.LittleEndian.AppendUint64([]byte(value), uint64(seed)))

	var x uint64
	for i := 0; i < 32; i += 8 {
		x ^= binary.LittleEndian.Uint64(h[i : i+8])
	}

	return uint(x)
}
