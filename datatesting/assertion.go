package datatesting

import "fmt"

func AssertEqual[T comparable](want, got T) error {
	if want != got {
		return fmt.Errorf("test failed: want %v, got %v", want, got)
	}

	return nil
}
