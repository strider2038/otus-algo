package structs_test

import (
	"fmt"
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/pkg/structs"
)

func TestBitSet_IsSet(t *testing.T) {
	bits := structs.NewBitSet(100)

	bits.Set(0)
	bits.Set(2)
	bits.Set(63)
	bits.Set(64)
	bits.Set(99)

	for i := 0; i <= 100; i++ {
		if i == 0 || i == 2 || i == 63 || i == 64 || i == 99 {
			datatesting.AssertTrue(t, bits.IsSet(i))
		} else {
			datatesting.AssertFalse(t, bits.IsSet(i))
		}
	}
}

func TestBitSet_OnesCount(t *testing.T) {
	bits := structs.NewBitSet(100)
	bits.Set(0)
	bits.Set(2)
	bits.Set(63)
	bits.Set(64)
	bits.Set(99)
	bits.Unset(2)

	count := bits.OnesCount()

	datatesting.AssertEqual(t, 4, count)
}

func TestBitSet_And(t *testing.T) {
	tests := []struct {
		bits1 structs.BitSet
		bits2 structs.BitSet
		want  structs.BitSet
	}{
		{bits1: structs.BitSet{0x11}, bits2: structs.BitSet{0xF}, want: structs.BitSet{0x1}},
		{bits1: structs.BitSet{0xF}, bits2: structs.BitSet{0xF0}, want: structs.BitSet{0x0}},
		{bits1: structs.BitSet{0x1}, bits2: structs.BitSet{0xF, 0xF}, want: structs.BitSet{0x1, 0x0}},
		{bits1: structs.BitSet{0x0, 0x1}, bits2: structs.BitSet{0xF}, want: structs.BitSet{0x0, 0x0}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s & %s", test.bits1, test.bits2), func(t *testing.T) {
			got := test.bits1.And(test.bits2)

			datatesting.AssertEqualArrays(t, test.want, got)
		})
	}
}

func TestBitSet_Or(t *testing.T) {
	tests := []struct {
		bits1 structs.BitSet
		bits2 structs.BitSet
		want  structs.BitSet
	}{
		{bits1: structs.BitSet{0x11}, bits2: structs.BitSet{0xF}, want: structs.BitSet{0x1F}},
		{bits1: structs.BitSet{0xF}, bits2: structs.BitSet{0xF0}, want: structs.BitSet{0xFF}},
		{bits1: structs.BitSet{0x1}, bits2: structs.BitSet{0xF, 0xF}, want: structs.BitSet{0xF, 0xF}},
		{bits1: structs.BitSet{0x0, 0x1}, bits2: structs.BitSet{0xF}, want: structs.BitSet{0xF, 0x1}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s & %s", test.bits1, test.bits2), func(t *testing.T) {
			got := test.bits1.Or(test.bits2)

			datatesting.AssertEqualArrays(t, test.want, got)
		})
	}
}

func TestBitSet_Intersects(t *testing.T) {
	tests := []struct {
		bits1 structs.BitSet
		bits2 structs.BitSet
		want  bool
	}{
		{bits1: structs.BitSet{0x1}, bits2: structs.BitSet{0xF}, want: true},
		{bits1: structs.BitSet{0xF}, bits2: structs.BitSet{0xF0}, want: false},
		{bits1: structs.BitSet{0x1}, bits2: structs.BitSet{0xF, 0xF}, want: true},
		{bits1: structs.BitSet{0x1}, bits2: structs.BitSet{0xF0, 0xF}, want: false},
		{bits1: structs.BitSet{0x1, 0x0}, bits2: structs.BitSet{0xF}, want: true},
		{bits1: structs.BitSet{0x0, 0x1}, bits2: structs.BitSet{0xF}, want: false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s & %s = %t", test.bits1, test.bits2, test.want), func(t *testing.T) {
			got := test.bits1.Intersects(test.bits2)

			datatesting.AssertEqual(t, test.want, got)
		})
	}
}
