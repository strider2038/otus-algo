package datatesting

import (
	"fmt"
	"math"
	"strings"
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
