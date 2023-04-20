package goroh_test

import (
	"testing"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-21-dynamic-programming/goroh"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{in: "3/5 + 7/9", want: "62/45"},
		{in: "3/13 + 7/26", want: "1/2"},
		{in: "9/21 + 5/7", want: "8/7"},
	}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			got, err := goroh.Calculate(test.in)

			if err != nil {
				t.Error(err)
				return
			}
			datatesting.AssertEqual(t, test.want, got)
		})
	}
}
