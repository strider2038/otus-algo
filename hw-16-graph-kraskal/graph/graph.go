package graph

import (
	"errors"

	"github.com/strider2038/otus-algo/hw-08-quick-external-sorts/sort"
	"github.com/strider2038/otus-algo/pkg/structs"
)

var ErrNonIntersectingGraphs = errors.New("non intersecting graphs")

type Vector []int

// AdjacencyArray - представление графа в виде массива векторов смежности.
type AdjacencyArray []Vector

type Edge struct {
	Vertex1 int
	Vertex2 int
	Weight  float64
}

type Edges []Edge

// FindSpanningTree находит остовное дерево с помощью алгоритма Краскала.
func (edges Edges) FindSpanningTree() (Edges, error) {
	if len(edges) == 0 {
		return edges, nil
	}

	sorted := edges.SortByWeight()
	verticesCount := edges.VerticesCount()

	// инициализация остовного дерева и списка множеств вершин на основе битовых масок
	spanningTree, spans := edges.initSpans(sorted[0], verticesCount)

	for i := 1; i < len(sorted) && len(spanningTree) < verticesCount-1; i++ {
		edge := sorted[i]
		hasCycle, hasIntersection := find(spans, edge)

		// если обнаружен цикл, то игнорируем ребро
		if hasCycle {
			continue
		}
		// добавляем ребро к остовному дереву
		spanningTree = append(spanningTree, edge)

		// если найдено пересечение, то объединяем множества остовов
		if hasIntersection {
			spans = union(spans)
		} else {
			// если пересечения не найдены, то добавляем еще одно множество остовов
			span := structs.NewBitSet(verticesCount)
			span.Set(edge.Vertex1)
			span.Set(edge.Vertex2)
			spans = append(spans, span)
		}
	}

	// если остовов больше одного, то граф непересекающийся
	if len(spans) > 1 {
		return nil, ErrNonIntersectingGraphs
	}

	return spanningTree, nil
}

// SortByWeight сортирует ребра по возрастанию веса. Возвращает копию исходного массива.
func (edges Edges) SortByWeight() Edges {
	sorted := make(Edges, len(edges))
	copy(sorted, edges)

	sort.QuickSlice(sorted, func(i, j int) bool {
		return sorted[i].Weight > sorted[j].Weight
	})

	return sorted
}

// VerticesCount подсчитывает число уникальных вершин в массиве ребер.
func (edges Edges) VerticesCount() int {
	vertices := structs.NewBitSet(len(edges))

	for _, edge := range edges {
		vertices.Set(edge.Vertex1)
		vertices.Set(edge.Vertex2)
	}

	return vertices.OnesCount()
}

func (edges Edges) initSpans(edge Edge, verticesCount int) (Edges, []structs.BitSet) {
	spanningTree := Edges{edge}
	firstSpan := structs.NewBitSet(verticesCount)
	firstSpan.Set(edge.Vertex1)
	firstSpan.Set(edge.Vertex2)
	spans := []structs.BitSet{firstSpan}

	return spanningTree, spans
}

func find(spans []structs.BitSet, edge Edge) (bool, bool) {
	cycleFound := false
	intersectionFound := false

	for _, span := range spans {
		// найден цикл
		if span.IsSet(edge.Vertex1) && span.IsSet(edge.Vertex2) {
			cycleFound = true
			break
		}

		// найдено пересечение по одной вершине
		if span.IsSet(edge.Vertex1) || span.IsSet(edge.Vertex2) {
			intersectionFound = true
			// объединяем два множества
			span.Set(edge.Vertex1)
			span.Set(edge.Vertex2)
			break
		}
	}

	return cycleFound, intersectionFound
}

func union(spans []structs.BitSet) []structs.BitSet {
	for i := 0; i < len(spans)-1; i++ {
		for j := i + 1; j < len(spans); {
			if spans[i].Intersects(spans[j]) {
				spans[i] = spans[i].Or(spans[j])
				// удаление элемента из слайса
				spans = append(spans[:j], spans[j+1:]...)
			} else {
				j++
			}
		}
	}

	return spans
}
