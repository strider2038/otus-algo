package datatesting

import (
	"math"
	"testing"
)

const delta = 0.000001

func AssertEqual[T comparable](t *testing.T, want, got T) {
	t.Helper()
	if want != got {
		t.Errorf("test failed: want %v, got %v", want, got)
	}
}

func AssertTrue(t *testing.T, isTrue bool) {
	t.Helper()
	if !isTrue {
		t.Errorf("test failed: want true, got false")
	}
}

func AssertFalse(t *testing.T, isTrue bool) {
	t.Helper()
	if isTrue {
		t.Errorf("test failed: want false, got true")
	}
}

func AssertEqualArrays[T comparable](t *testing.T, wantItems []T, got []T) {
	t.Helper()

	if len(wantItems) != len(got) {
		t.Errorf("different length: want %d, got %d", len(wantItems), len(got))
		return
	}

	errsCount := 0

	for i := 0; i < len(wantItems); i++ {
		if wantItems[i] != got[i] {
			t.Errorf("different items at %d: want %v, got %v", i, wantItems[i], got[i])
			errsCount++
			if errsCount > 100 {
				t.Errorf("too much errors")
				return
			}
		}
	}
}

func AssertEqualMatrix(t *testing.T, want [][]int, got [][]int) {
	t.Helper()

	if len(want) != len(got) {
		t.Errorf("different matrix rows count: want %d, got %d", len(want), len(got))
		return
	}

	for i := 0; i < len(want); i++ {
		wantRow := want[i]
		gotRow := got[i]

		if len(wantRow) != len(gotRow) {
			t.Errorf("different length at row %d: want %d, got %d", i, len(wantRow), len(gotRow))
			continue
		}

		errsCount := 0
		for j := 0; j < len(wantRow); j++ {
			if wantRow[j] != gotRow[j] {
				t.Errorf("different items at [%d][%d]: want %d, got %d", i, j, wantRow[j], gotRow[j])
				errsCount++
				if errsCount > 100 {
					t.Errorf("too much errors")
					return
				}
			}
		}
	}
}

func AssertEqualFloat(t *testing.T, want, got float64) {
	t.Helper()
	AssertEqualFloatWithDelta(t, want, got, delta)
}

func AssertEqualFloatWithDelta(t *testing.T, want, got, delta float64) {
	if math.Abs(want-got) > delta {
		t.Errorf("test failed: want %v, got %v", want, got)
	}
}
