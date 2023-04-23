package obst_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-11-optimal-search-trees/obst"
)

func TestNewV1(t *testing.T) {
	tree := obst.NewV1(
		obst.InputNode[int]{Weight: 90, Value: 1},
		obst.InputNode[int]{Weight: 10, Value: 2},
		obst.InputNode[int]{Weight: 30, Value: 3},
		obst.InputNode[int]{Weight: 60, Value: 4},
		obst.InputNode[int]{Weight: 40, Value: 5},
		obst.InputNode[int]{Weight: 50, Value: 6},
		obst.InputNode[int]{Weight: 20, Value: 7},
		obst.InputNode[int]{Weight: 95, Value: 8},
		obst.InputNode[int]{Weight: 15, Value: 9},
	)

	values := make([]int, 0)
	weights := make([]float64, 0)
	_ = tree.ForEach(func(value int, weight float64) error {
		values = append(values, value)
		weights = append(weights, weight)
		return nil
	})

	datatesting.AssertEqualArrays(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, values)
	datatesting.AssertEqualArrays(t, []float64{90, 10, 30, 60, 40, 50, 20, 95, 15}, weights)
}

func TestNewV2(t *testing.T) {
	tree := obst.NewV2(
		obst.InputNode[int]{Weight: 90, Value: 1},
		obst.InputNode[int]{Weight: 10, Value: 2},
		obst.InputNode[int]{Weight: 30, Value: 3},
		obst.InputNode[int]{Weight: 60, Value: 4},
		obst.InputNode[int]{Weight: 40, Value: 5},
		obst.InputNode[int]{Weight: 50, Value: 6},
		obst.InputNode[int]{Weight: 20, Value: 7},
		obst.InputNode[int]{Weight: 95, Value: 8},
		obst.InputNode[int]{Weight: 15, Value: 9},
	)

	values := make([]int, 0)
	weights := make([]float64, 0)
	_ = tree.ForEach(func(value int, weight float64) error {
		values = append(values, value)
		weights = append(weights, weight)
		return nil
	})

	datatesting.AssertEqualArrays(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, values)
	datatesting.AssertEqualArrays(t, []float64{90, 10, 30, 60, 40, 50, 20, 95, 15}, weights)
}

func TestPerformance(t *testing.T) {
	algos := []struct {
		name   string
		create func(nodes ...obst.InputNode[int]) *obst.Tree[int]
	}{
		{
			name:   "algo v1",
			create: obst.NewV1[int],
		},
		{
			name:   "algo v2",
			create: obst.NewV2[int],
		},
	}

	sizes := []int{100, 1_000, 10_000, 100_000}

	for _, size := range sizes {
		nodes := createNodes(size)

		for _, algo := range algos {
			t.Run(fmt.Sprintf("%s, n = %d", algo.name, size), func(t *testing.T) {
				start := time.Now()
				tree := algo.create(nodes...)
				t.Log("tree creation time:", time.Since(start))

				t.Log("search min weighted value:", testSearchTime(t, tree, 1, 1000))
				t.Log("search mid weighted value:", testSearchTime(t, tree, size/2+size/4, 1000))
				t.Log("search max weighted value:", testSearchTime(t, tree, size-1, 1000))
				t.Log("search random value:", testRandomSearchTime(t, tree, size, 1000))
			})
		}
	}
}

func createNodes(count int) []obst.InputNode[int] {
	nodes := make([]obst.InputNode[int], count)

	// создание массива узлов с равномерно возрастающими значениями
	// и весами, возрастающими по экспоненте
	for i := 0; i < count; i++ {
		nodes[i].Value = i + 1
		nodes[i].Weight = math.Exp(3*float64(i)/float64(count-1) - 3)
	}

	// перемешивание массива случайным образом
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(nodes), func(i, j int) {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	})

	return nodes
}

func testSearchTime(t *testing.T, tree *obst.Tree[int], value int, n int) time.Duration {
	duration := time.Duration(0)

	for i := 0; i < n; i++ {
		start := time.Now()
		_, err := tree.Find(value)
		if err != nil {
			t.Fatalf("search for %v: %v", value, err)
		}
		duration += time.Since(start)
	}

	return duration
}

func testRandomSearchTime(t *testing.T, tree *obst.Tree[int], count int, n int) time.Duration {
	duration := time.Duration(0)

	for i := 0; i < n; i++ {
		value := rand.Intn(count-1) + 1
		start := time.Now()
		_, err := tree.Find(value)
		if err != nil {
			t.Fatalf("search for %v: %v", value, err)
		}
		duration += time.Since(start)
	}

	return duration
}
