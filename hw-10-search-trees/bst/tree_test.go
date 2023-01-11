package bst_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/strider2038/otus-algo/datatesting"
	"github.com/strider2038/otus-algo/hw-06-basic-sorts/basicsorts"
	"github.com/strider2038/otus-algo/hw-10-search-trees/bst"
)

func TestTree_Insert(t *testing.T) {
	tree := bst.Tree[int]{}

	tree.Insert(5)
	tree.Insert(8)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(11)
	tree.Insert(7)

	got := make([]int, 0)
	tree.ForEach(func(value int) error {
		got = append(got, value)
		return nil
	})
	datatesting.AssertEqualArrays(t, []int{1, 3, 5, 7, 8, 11}, got)
}

func TestTree_Remove(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		remove int
		want   []int
	}{
		{
			name:   "remove not existing node",
			values: []int{10, 3, 5, 15, 13, 11},
			remove: -1,
			want:   []int{3, 5, 10, 11, 13, 15},
		},
		{
			name:   "remove root node",
			values: []int{2, 1, 3},
			remove: 2,
			want:   []int{1, 3},
		},
		{
			name:   "remove left node without children",
			values: []int{2, 1, 3},
			remove: 1,
			want:   []int{2, 3},
		},
		{
			name:   "remove right node without children",
			values: []int{2, 1, 3},
			remove: 3,
			want:   []int{1, 2},
		},
		{
			name:   "remove node with right child",
			values: []int{4, 2, 3},
			remove: 2,
			want:   []int{3, 4},
		},
		{
			name:   "remove node with left child",
			values: []int{4, 2, 1},
			remove: 2,
			want:   []int{1, 4},
		},
		{
			name:   "remove node with both children",
			values: []int{5, 2, 4, 1, 3},
			remove: 2,
			want:   []int{1, 3, 4, 5},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tree := createTreeFromNumbers(test.values)

			tree.Remove(test.remove)

			got := make([]int, 0)
			tree.ForEach(func(value int) error {
				got = append(got, value)
				return nil
			})
			datatesting.AssertEqualArrays(t, test.want, got)
		})
	}
}

func TestTree(t *testing.T) {
	treeOperations := []struct {
		name string
		do   func(tree *bst.Tree[int], values []int) error
	}{
		{
			name: "find",
			do: func(tree *bst.Tree[int], values []int) error {
				value := values[rand.Intn(len(values))]
				_, err := tree.Find(value)
				return err
			},
		},
		{
			name: "remove",
			do: func(tree *bst.Tree[int], values []int) error {
				tree.Remove(values[rand.Intn(len(values))])
				return nil
			},
		},
	}

	counts := []int{1000, 10_000, 100_000}

	for _, operation := range treeOperations {
		for _, count := range counts {
			numbers := GenerateRandomNumbers(count, count)

			t.Run(fmt.Sprintf("%s, n=%d, random", operation.name, count), func(t *testing.T) {
				tree := createTreeFromNumbers(numbers)

				start := time.Now()
				for i := 0; i < count/10; i++ {
					if err := operation.do(tree, numbers); err != nil {
						t.Fatal(err)
					}
				}
				t.Log("elapsed time:", time.Since(start).String())
			})

			numbers = basicsorts.Shell(numbers)
			t.Run(fmt.Sprintf("%s, n=%d, sorted", operation.name, count), func(t *testing.T) {
				tree := createTreeFromNumbers(numbers)

				start := time.Now()
				for i := 0; i < count/10; i++ {
					if err := operation.do(tree, numbers); err != nil {
						t.Error(err)
					}
				}
				t.Log("elapsed time:", time.Since(start).String())
			})
		}
	}
}

func GenerateRandomNumbers(n int, max int) []int {
	rand.Seed(time.Now().UnixNano())

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(max)
	}

	return numbers
}

func createTreeFromNumbers(values []int) *bst.Tree[int] {
	tree := &bst.Tree[int]{}
	for _, value := range values {
		tree.Insert(value)
	}
	return tree
}
