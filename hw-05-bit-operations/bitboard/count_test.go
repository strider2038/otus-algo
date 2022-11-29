package bitboard_test

import (
	"fmt"
	"math/bits"
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
		{bits: 0xffffffffffffffff, want: 64},
	}
	counters := []struct {
		name  string
		count func(uint64) int
	}{
		{name: "sequentially", count: bitboard.CountBitsSequentially},
		{name: "by division", count: bitboard.CountBitsByDivision},
		{name: "by 8-bit cache", count: bitboard.CountBitsByCache},
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

// cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
// BenchmarkCountBitsSequentially
// BenchmarkCountBitsSequentially-16    	15998384	        74.35 ns/op
func BenchmarkCountBitsSequentially(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitboard.CountBitsSequentially(0x1)
		bitboard.CountBitsSequentially(0x8000000000000000)
		bitboard.CountBitsSequentially(0x0123456789abcdef)
		bitboard.CountBitsSequentially(0xfedcba9876543210)
		bitboard.CountBitsSequentially(0xffffffffffffffff)
	}
}

// cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
// BenchmarkCountBitsByDivision
// BenchmarkCountBitsByDivision-16    	24209925	        50.24 ns/op
func BenchmarkCountBitsByDivision(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitboard.CountBitsByDivision(0x1)
		bitboard.CountBitsByDivision(0x8000000000000000)
		bitboard.CountBitsByDivision(0x0123456789abcdef)
		bitboard.CountBitsByDivision(0xfedcba9876543210)
		bitboard.CountBitsByDivision(0xffffffffffffffff)
	}
}

// cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
// BenchmarkCountBitsByCache
// BenchmarkCountBitsByCache-16    	69435896	        16.50 ns/op
func BenchmarkCountBitsByCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitboard.CountBitsByCache(0x1)
		bitboard.CountBitsByCache(0x8000000000000000)
		bitboard.CountBitsByCache(0x0123456789abcdef)
		bitboard.CountBitsByCache(0xfedcba9876543210)
		bitboard.CountBitsByCache(0xffffffffffffffff)
	}
}

// cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
// BenchmarkCountBitsByStdLib
// BenchmarkCountBitsByStdLib-16    	909551740	         1.265 ns/op
func BenchmarkCountBitsByStdLib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bits.OnesCount(0x1)
		bits.OnesCount(0x8000000000000000)
		bits.OnesCount(0x0123456789abcdef)
		bits.OnesCount(0xfedcba9876543210)
		bits.OnesCount(0xffffffffffffffff)
	}
}
