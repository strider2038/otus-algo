package rle2_test

import (
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-20-run-length-encoding/rle2"
)

func FuzzCompress(f *testing.F) {
	f.Add([]byte{1})
	f.Fuzz(func(t *testing.T, raw []byte) {
		compressed := rle2.CompressBytes(raw)
		decompressed := rle2.DecompressBytes(compressed)
		datatesting.AssertEqualArrays(t, raw, decompressed)
	})
}
