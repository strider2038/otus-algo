package arrays_test

import (
	"testing"

	"github.com/strider2038/otus-algo/hw-04-basic-data-structures/arrays"
)

func TestArrays(t *testing.T) {
	tests := []struct {
		name      string
		array     []int
		operation func(arr arrays.Array[int]) error
		want      []int
	}{
		{
			name:      "insert into middle",
			array:     []int{1, 2, 3, 4, 5},
			operation: func(arr arrays.Array[int]) error { return arr.Insert(3, 100) },
			want:      []int{1, 2, 3, 100, 4, 5},
		},
		{
			name:      "insert first",
			array:     []int{1, 2, 3, 4, 5},
			operation: func(arr arrays.Array[int]) error { return arr.Insert(0, 100) },
			want:      []int{100, 1, 2, 3, 4, 5},
		},
		{
			name:      "insert last",
			array:     []int{1, 2, 3, 4, 5},
			operation: func(arr arrays.Array[int]) error { return arr.Insert(5, 100) },
			want:      []int{1, 2, 3, 4, 5, 100},
		},
		{
			name:      "insert into empty array",
			array:     []int{},
			operation: func(arr arrays.Array[int]) error { return arr.Insert(0, 100) },
			want:      []int{100},
		},
		{
			name:      "remove from middle",
			array:     []int{1, 2, 3, 4, 5},
			operation: remove(2),
			want:      []int{1, 2, 4, 5},
		},
		{
			name:      "remove first",
			array:     []int{1, 2, 3, 4, 5},
			operation: remove(0),
			want:      []int{2, 3, 4, 5},
		},
		{
			name:      "remove last",
			array:     []int{1, 2, 3, 4, 5},
			operation: remove(4),
			want:      []int{1, 2, 3, 4},
		},
		{
			name:      "remove from single array",
			array:     []int{1},
			operation: remove(0),
			want:      []int{},
		},
	}
	arrayTypes := []struct {
		name   string
		create func() arrays.Array[int]
	}{
		{"SingleArray", func() arrays.Array[int] { return arrays.NewSingleArray[int]() }},
		{"VectorArray(3)", func() arrays.Array[int] { return arrays.NewVectorArray[int](3) }},
		{"VectorArray(5)", func() arrays.Array[int] { return arrays.NewVectorArray[int](5) }},
		{"FactorArray(2, 0)", func() arrays.Array[int] { return arrays.NewFactorArray[int](2, 0) }},
		{"FactorArray(5, 0)", func() arrays.Array[int] { return arrays.NewFactorArray[int](5, 0) }},
		{"FactorArray(2, 5)", func() arrays.Array[int] { return arrays.NewFactorArray[int](2, 5) }},
		{"MatrixArray(3)", func() arrays.Array[int] { return arrays.NewMatrixArray[int](3) }},
		{"MatrixArray(5)", func() arrays.Array[int] { return arrays.NewMatrixArray[int](5) }},
		{"SliceArray", func() arrays.Array[int] { return arrays.NewSliceArray[int]() }},
	}
	for _, arrayType := range arrayTypes {
		for _, test := range tests {
			t.Run(arrayType.name+": "+test.name, func(t *testing.T) {
				array := arrayType.create()
				for _, element := range test.array {
					array.Add(element)
				}

				err := test.operation(array)

				if err != nil {
					t.Fatal("operation:", err)
				}
				if array.Size() != len(test.want) {
					t.Fatalf("array size mismatch: want %d, got %d", len(test.want), array.Size())
				}
				for i := 0; i < len(test.want); i++ {
					got, err := array.Get(i)
					if err != nil {
						t.Fatalf("get array element by index %d: %s", i, err)
					}
					if got != test.want[i] {
						t.Errorf("unexpected array element at %d: want %d, got %d", i, test.want[i], got)
					}
				}
			})
		}
	}
}

func remove(index int) func(arr arrays.Array[int]) error {
	return func(arr arrays.Array[int]) error {
		_, err := arr.Remove(index)
		return err
	}
}
