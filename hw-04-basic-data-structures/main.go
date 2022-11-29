package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/strider2038/otus-algo/hw-04-basic-data-structures/arrays"
)

type Array[T any] interface {
	Size() int
	Get(index int) (T, error)
	Add(item T)
	Insert(index int, item T) error
	Remove(index int) (T, error)
}

func main() {
	RunArrayTest(
		"SingleArray()",
		func() Array[int] {
			return arrays.NewSingleArray[int]()
		},
		4,
	)
	RunArrayTest(
		"VectorArray(100)",
		func() Array[int] {
			return arrays.NewVectorArray[int](100)
		},
		4,
	)
	RunArrayTest(
		"FactorArray(2)",
		func() Array[int] {
			return arrays.NewFactorArray[int](2, 0)
		},
		4,
	)
	RunArrayTest(
		"MatrixArray(100)",
		func() Array[int] {
			return arrays.NewMatrixArray[int](100)
		},
		3,
	)
	RunArrayTest(
		"SliceArray()",
		func() Array[int] { return arrays.NewSliceArray[int]() },
		4,
	)
}

func RunArrayTest(arrayType string, newArray func() Array[int], n int) {
	tester := &ArrayTester{
		ArrayType: arrayType,
		NewArray:  newArray,
	}

	count := 100
	for i := 0; i < n; i++ {
		tester.Counts = append(tester.Counts, count)
		count *= 10
	}

	tester.Run()
}

type ArrayTester struct {
	ArrayType string
	NewArray  func() Array[int]
	Counts    []int
}

func (t *ArrayTester) Run() {
	t.testEmptyArray(
		"Add",
		func(arr Array[int]) { arr.Add(rand.Int()) },
	)
	t.testEmptyArray(
		"Insert(0)",
		func(arr Array[int]) { arr.Insert(0, rand.Int()) },
	)
	t.testEmptyArray(
		"Insert(n/2)",
		func(arr Array[int]) { arr.Insert(arr.Size()/2, rand.Int()) },
	)
	t.testFilledArray(
		"Remove(0)",
		func(arr Array[int]) { arr.Remove(0) },
	)
	t.testFilledArray(
		"Remove(n/2)",
		func(arr Array[int]) { arr.Remove(arr.Size() / 2) },
	)
	t.testFilledArray(
		"Remove(n-1)",
		func(arr Array[int]) { arr.Remove(arr.Size() - 1) },
	)
}

func (t *ArrayTester) testEmptyArray(method string, test func(arr Array[int])) {
	for _, count := range t.Counts {
		array := t.NewArray()
		t.testMethod(method, array, test, count)
	}
}

func (t *ArrayTester) testFilledArray(method string, test func(arr Array[int])) {
	for _, count := range t.Counts {
		array := t.NewArray()
		for i := 0; i < count; i++ {
			array.Add(rand.Int())
		}
		t.testMethod(method, array, test, count)
	}
}

func (t *ArrayTester) testMethod(
	method string,
	array Array[int],
	test func(arr Array[int]),
	count int,
) {
	// принудительный запуск сборщика мусора
	runtime.GC()
	start := time.Now()
	for i := 0; i < count; i++ {
		test(array)
	}
	elapsed := time.Since(start)
	fmt.Printf("%s.%s: n=%d, elapsed=%s\n", t.ArrayType, method, count, elapsed.String())
}
