package datatesting

import (
	"fmt"
	"math"
)

const delta = 0.000001

func AssertEqual[T comparable](want, got T) error {
	if want != got {
		return fmt.Errorf("test failed: want %v, got %v", want, got)
	}

	return nil
}

func AssertEqualFloat(want, got float64) error {
	return AssertEqualFloatWithDelta(want, got, delta)
}

func AssertEqualFloatWithDelta(want, got, delta float64) error {
	if math.Abs(want-got) > delta {
		return fmt.Errorf("test failed: want %v, got %v", want, got)
	}

	return nil
}
