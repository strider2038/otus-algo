package bitboard_test

import (
	"fmt"
	"testing"

	"github.com/strider2038/otus-algo/hw-05-bit-operations/bitboard"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		bits uint64
		want int
	}{
		{bits: 0x1, want: 1},
		{bits: 0xF, want: 4},
		{bits: 0x1234, want: 5},
		{bits: 0x123456789abcdef, want: 32},
	}
	counters := []struct {
		name  string
		count func(uint64) int
	}{
		{name: "sequentially", count: bitboard.CountBitsSequentially},
		{name: "by division", count: bitboard.CountBitsByDivision},
	}
	for _, counter := range counters {
		for _, test := range tests {
			t.Run(fmt.Sprintf("%s: %x %d", counter.name, test.bits, test.want), func(t *testing.T) {
				got := counter.count(test.bits)

				if test.want != got {
					t.Errorf("want %d, got %d", test.want, got)
				}
			})
		}
	}
}
