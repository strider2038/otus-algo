package bitboard_test

import (
	"strconv"
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-05-bit-operations/bitboard"
)

type Solver func(board uint8) (int, uint64)

func (s Solver) Solve(t *testing.T, input, output []string) {
	if len(input) < 1 || len(output) < 2 {
		t.Fatal(datatesting.ErrNotEnoughArguments)
	}

	board, err := strconv.ParseUint(input[0], 10, 8)
	if err != nil {
		t.Fatalf("parse input: %v", err)
	}
	wantCount, err := strconv.ParseInt(output[0], 10, 64)
	if err != nil {
		t.Fatalf("parse output[0]: %v", err)
	}
	wantMoves, err := strconv.ParseUint(output[1], 10, 64)
	if err != nil {
		t.Fatalf("parse output[1]: %v", err)
	}

	gotCount, gotMoves := s(uint8(board))

	datatesting.AssertEqual(t, wantCount, int64(gotCount))
	datatesting.AssertEqual(t, wantMoves, gotMoves)
}

func TestKingMoves(t *testing.T) {
	runner := datatesting.NewRunner(
		datatesting.WithWorkdir("./testdata/1_king/"),
		datatesting.WithSeparator("\n"),
	)
	runner.Run(t, Solver(bitboard.KingMoves))
}

func TestKnightMoves(t *testing.T) {
	runner := datatesting.NewRunner(
		datatesting.WithWorkdir("./testdata/2_knight/"),
		datatesting.WithSeparator("\n"),
	)
	runner.Run(t, Solver(bitboard.KnightMoves))
}

func TestRookMoves(t *testing.T) {
	runner := datatesting.NewRunner(
		datatesting.WithWorkdir("./testdata/3_rook/"),
		datatesting.WithSeparator("\n"),
	)
	runner.Run(t, Solver(bitboard.RookMoves))
}
