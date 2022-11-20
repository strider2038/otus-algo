package bitboard_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-05-bit-operations/bitboard"
)

type Solver func(board uint8) (int, uint64)

func (s Solver) Solve(input, output []string) error {
	if len(input) < 1 || len(output) < 2 {
		return datatesting.ErrNotEnoughArguments
	}

	board, err := strconv.ParseUint(input[0], 10, 8)
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}
	wantCount, err := strconv.ParseInt(output[0], 10, 64)
	if err != nil {
		return fmt.Errorf("parse output[0]: %w", err)
	}
	wantMoves, err := strconv.ParseUint(output[1], 10, 64)
	if err != nil {
		return fmt.Errorf("parse output[1]: %w", err)
	}

	gotCount, gotMoves := s(uint8(board))

	return datatesting.AssertNoErrors(
		datatesting.AssertEqual(wantCount, int64(gotCount)),
		datatesting.AssertEqual(wantMoves, gotMoves),
	)
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
