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

func TestCompressBytes(t *testing.T) {
	for i, test := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			compressed := rle.CompressBytes(test.decompressed)

			datatesting.AssertEqualArrays(t, test.compressed, compressed)
		})
	}
}

func TestDecompressBytes(t *testing.T) {
	for i, test := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			decompressed := rle.DecompressBytes(test.compressed)

			datatesting.AssertEqualArrays(t, test.decompressed, decompressed)
		})
	}
}

func TestCompress(t *testing.T) {
	for i, test := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			compressed := bytes.Buffer{}

			err := rle.Compress(bytes.NewReader(test.decompressed), &compressed)

			if err != nil {
				t.Fatal(err)
			}
			datatesting.AssertEqualArrays(t, test.compressed, compressed.Bytes())
		})
	}
}

func TestDecompress(t *testing.T) {
	for i, test := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			decompressed := bytes.Buffer{}

			err := rle.Decompress(bytes.NewReader(test.compressed), &decompressed)

			if err != nil {
				t.Fatal(err)
			}
			datatesting.AssertEqualArrays(t, test.decompressed, decompressed.Bytes())
		})
	}
}
