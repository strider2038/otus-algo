package datatesting

import (
	"math"
	"strings"
	"testing"
)

const delta = 0.000001

func AssertEqual[T comparable](t *testing.T, want, got T) {
	t.Helper()
	if want != got {
		t.Errorf("test failed: want %v, got %v", want, got)
	}
}

func AssertEqualArrays(t *testing.T, wantItems []int, got []int) {
	t.Helper()

	if len(wantItems) != len(got) {
		t.Errorf("different length: want %d, got %d", len(wantItems), len(got))
		return
	}

	errsCount := 0

	for i := 0; i < len(wantItems); i++ {
		if wantItems[i] != got[i] {
			t.Errorf("different items at %d: want %d, got %d", i, wantItems[i], got[i])
			errsCount++
			if errsCount > 100 {
				t.Errorf("too much errors")
				return
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

func AssertNoErrors(errs ...error) error {
	var s strings.Builder

	for _, err := range errs {
		if err != nil {
			if s.Len() > 0 {
				s.WriteString("; ")
			}
			s.WriteString(err.Error())
		}
	}

	if s.Len() > 0 {
		return fmt.Errorf(s.String())
	}

	return nil
}
