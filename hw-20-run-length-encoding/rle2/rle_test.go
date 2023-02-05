package rle2_test

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-20-run-length-encoding/rle2"
)

var cases = []struct {
	decompressed []byte
	compressed   []byte
}{
	{
		decompressed: nil,
		compressed:   nil,
	},
	{
		decompressed: []byte{1},
		compressed:   []byte{1, 1},
	},
	{
		decompressed: []byte{1, 2},
		compressed:   []byte{254, 1, 2},
	},
	{
		decompressed: []byte{1, 1, 2},
		compressed:   []byte{2, 1, 1, 2},
	},
	{
		decompressed: []byte{1, 1, 2, 3},
		compressed:   []byte{2, 1, 254, 2, 3},
	},
	{
		decompressed: []byte{1, 2, 3, 4, 5, 6, 7, 8},
		compressed:   []byte{248, 1, 2, 3, 4, 5, 6, 7, 8},
	},
	{
		decompressed: []byte{1, 2, 3, 4, 5, 5},
		compressed:   []byte{252, 1, 2, 3, 4, 2, 5},
	},
	{
		decompressed: []byte{1, 2, 3, 4, 5, 5, 5, 5},
		compressed:   []byte{252, 1, 2, 3, 4, 4, 5},
	},
	{
		decompressed: []byte{1, 1, 1, 2, 2, 2, 2, 3, 3, 4, 5},
		compressed:   []byte{3, 1, 4, 2, 2, 3, 254, 4, 5},
	},
	{
		decompressed: []byte{1, 2, 2, 3, 3},
		compressed:   []byte{255, 1, 2, 2, 2, 3},
	},
	{
		decompressed: bytes.Repeat([]byte{1}, 300),
		compressed:   []byte{127, 1, 127, 1, 46, 1},
	},
	{
		decompressed: bytes.Repeat([]byte{1}, 128),
		compressed:   []byte{127, 1, 1, 1},
	},
	{
		decompressed: bytesSequence(0, 255),
		compressed: mergeBytes(
			[]byte{128}, bytesSequence(0, 127),
			[]byte{128}, bytesSequence(128, 255),
		),
	},
	{
		decompressed: mergeBytes(bytesSequence(0, 255), bytesSequence(0, 10)),
		compressed: mergeBytes(
			[]byte{128}, bytesSequence(0, 127),
			[]byte{128}, bytesSequence(128, 255),
			[]byte{245}, bytesSequence(0, 10),
		),
	},
	{
		decompressed: mergeBytes(bytesSequence(0, 255), bytes.Repeat([]byte{0}, 10)),
		compressed: mergeBytes(
			[]byte{128}, bytesSequence(0, 127),
			[]byte{128}, bytesSequence(128, 255),
			[]byte{10, 0},
		),
	},
}

func TestCompress(t *testing.T) {
	for i, test := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			compressed := rle2.Compress(test.decompressed)

			datatesting.AssertEqualArrays(t, test.compressed, compressed)
		})
	}
}

func TestDecompress(t *testing.T) {
	for i, test := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			decompressed := rle2.Decompress(test.compressed)

			datatesting.AssertEqualArrays(t, test.decompressed, decompressed)
		})
	}
}

func bytesSequence(from, to byte) []byte {
	b := make([]byte, int(to)-int(from)+1)
	for i := 0; i <= int(to)-int(from); i++ {
		b[i] = byte(i) + from
	}
	return b
}

func mergeBytes(bb ...[]byte) []byte {
	var m []byte

	for _, b := range bb {
		m = append(m, b...)
	}

	return m
}
