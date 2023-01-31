package graph

import (
	"errors"
	"math"

	"github.com/strider2038/otus-algo/pkg/structs"
)

var ErrPathNotFound = errors.New("path not found")

type Direction struct {
	Vertex   int
	Distance float64
}

type Vector []Direction

// AdjacencyArray - представление графа в виде массива векторов смежности.
type AdjacencyArray []Vector

type Edge struct {
	Vertex1 int
	Vertex2 int
}

// FindShortestPath ищет наиболее короткий путь от вершины from к вершине to с помощью
// алгоритма Дийкстры. Если путь не может быть найден, то возвращается ошибка ErrPathNotFound.
func (graph AdjacencyArray) FindShortestPath(from, to int) ([]Edge, error) {
	cameFrom := make([]int, len(graph))
	visited := structs.NewBitSet(len(graph))
	distances := make([]float64, len(graph))
	for i := range distances {
		distances[i] = math.Inf(1)
	}
	distances[from] = 0

	heap := newMinHeap(len(graph))
	heap.Insert(&Direction{Vertex: from, Distance: distances[from]})
	for {
		current, ok := heap.Pop()
		if !ok || current.Vertex == to {
			break
		}
		if visited.IsSet(current.Vertex) {
			continue
		}
		visited.Set(current.Vertex)

		for _, direction := range graph[current.Vertex] {
			if distances[direction.Vertex] > distances[current.Vertex]+direction.Distance {
				distances[direction.Vertex] = distances[current.Vertex] + direction.Distance
				cameFrom[direction.Vertex] = current.Vertex
				heap.Insert(&Direction{Vertex: direction.Vertex, Distance: direction.Distance})
			}
		}
	}

	if math.IsInf(distances[to], 1) {
		return nil, ErrPathNotFound
	}

	// проход по обратному пути с построением списка ребер
	path := make([]Edge, 0)
	for last := to; last != from; {
		previous := cameFrom[last]
		path = append([]Edge{{Vertex1: previous, Vertex2: last}}, path...)
		last = previous
	}

	return path, nil
}
