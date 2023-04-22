package bloomsha

import "github.com/strider2038/otus-algo/pkg/structs"

type Filter struct {
	m    uint
	k    uint
	bits structs.BitSet
}

func NewFilter(m uint, k uint) *Filter {
	return &Filter{m: m, k: k, bits: structs.NewBitSet(int(m))}
}
