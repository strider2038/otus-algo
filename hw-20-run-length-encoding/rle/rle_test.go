package rle_test

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-20-run-length-encoding/rle"
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
		compressed:   []byte{1, 1, 1, 2},
	},
	{
		decompressed: []byte{1, 1, 2},
		compressed:   []byte{2, 1, 1, 2},
	},
	{
		decompressed: []byte{1, 1, 1, 2, 2, 2, 2, 3, 3, 4, 5},
		compressed:   []byte{3, 1, 4, 2, 2, 3, 1, 4, 1, 5},
	},
	{
		decompressed: bytes.Repeat([]byte{1}, 300),
		compressed:   []byte{255, 1, 45, 1},
	},
}

func TestCompress(t *testing.T) {
	for i, test := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			compressed := rle.Compress(test.decompressed)

			datatesting.AssertEqualArrays(t, test.compressed, compressed)
		})
	}
}

func TestDecompress(t *testing.T) {
	for i, test := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			decompressed := rle.Decompress(test.compressed)

			datatesting.AssertEqualArrays(t, test.decompressed, decompressed)
		})
	}
}
