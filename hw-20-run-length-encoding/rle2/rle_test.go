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
	{ // 0
		decompressed: nil,
		compressed:   nil,
	},
	{ // 1
		decompressed: []byte{1},
		compressed:   []byte{1, 1},
	},
	{ // 2
		decompressed: []byte{1, 2},
		compressed:   []byte{254, 1, 2},
	},
	{ // 3
		decompressed: []byte{1, 1, 2},
		compressed:   []byte{2, 1, 1, 2},
	},
	{ // 4
		decompressed: []byte{1, 1, 2, 3},
		compressed:   []byte{2, 1, 254, 2, 3},
	},
	{ // 5
		decompressed: []byte{1, 2, 3, 4, 5, 6, 7, 8},
		compressed:   []byte{248, 1, 2, 3, 4, 5, 6, 7, 8},
	},
	{ // 6
		decompressed: []byte{1, 2, 3, 4, 5, 5},
		compressed:   []byte{252, 1, 2, 3, 4, 2, 5},
	},
	{ // 7
		decompressed: []byte{1, 2, 3, 4, 5, 5, 5, 5},
		compressed:   []byte{252, 1, 2, 3, 4, 4, 5},
	},
	{ // 8
		decompressed: []byte{1, 1, 1, 2, 2, 2, 2, 3, 3, 4, 5},
		compressed:   []byte{3, 1, 4, 2, 2, 3, 254, 4, 5},
	},
	{ // 9
		decompressed: []byte{1, 2, 2, 3, 3},
		compressed:   []byte{1, 1, 2, 2, 2, 3},
	},
	{ // 10
		decompressed: []byte{1, 1, 2, 3, 3},
		compressed:   []byte{2, 1, 1, 2, 2, 3},
	},
	{ // 11
		decompressed: bytes.Repeat([]byte{1}, 300),
		compressed:   []byte{127, 1, 127, 1, 46, 1},
	},
	{ // 12
		decompressed: bytes.Repeat([]byte{1}, 128),
		compressed:   []byte{127, 1, 1, 1},
	},
	{ // 13
		decompressed: bytesSequence(0, 255),
		compressed: mergeBytes(
			[]byte{128}, bytesSequence(0, 127),
			[]byte{128}, bytesSequence(128, 255),
		),
	},
	{ // 14
		decompressed: mergeBytes(bytesSequence(0, 255), bytesSequence(0, 10)),
		compressed: mergeBytes(
			[]byte{128}, bytesSequence(0, 127),
			[]byte{128}, bytesSequence(128, 255),
			[]byte{245}, bytesSequence(0, 10),
		),
	},
	{ // 15
		decompressed: mergeBytes(bytesSequence(0, 255), bytes.Repeat([]byte{0}, 10)),
		compressed: mergeBytes(
			[]byte{128}, bytesSequence(0, 127),
			[]byte{128}, bytesSequence(128, 255),
			[]byte{10, 0},
		),
	},
}

func TestCompressBytes(t *testing.T) {
	for i, test := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			compressed := rle2.CompressBytes(test.decompressed)

			datatesting.AssertEqualArrays(t, test.compressed, compressed)
		})
	}
}

func TestDecompressBytes(t *testing.T) {
	for i, test := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			decompressed := rle2.DecompressBytes(test.compressed)

			datatesting.AssertEqualArrays(t, test.decompressed, decompressed)
		})
	}
}

func TestCompress(t *testing.T) {
	for i, test := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			compressed := bytes.Buffer{}

			err := rle2.Compress(bytes.NewReader(test.decompressed), &compressed)

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

			err := rle2.Decompress(bytes.NewReader(test.compressed), &decompressed)

			if err != nil {
				t.Fatal(err)
			}
			datatesting.AssertEqualArrays(t, test.decompressed, decompressed.Bytes())
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
