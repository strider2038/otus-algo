package graph

import "github.com/strider2038/otus-algo/pkg/structs"

type Vector []int

type AdjacencyMatrix []Vector

// FindStronglyConnectedComponents - находит компоненты сильной связанности
// с помощью алгоритма Косарайю (https://habr.com/ru/post/537290/).
func (graph AdjacencyMatrix) FindStronglyConnectedComponents() [][]int {
	return newComponentsSeeker(graph).Find()
}

// Inverse - возвращает инвертированный граф.
func (graph AdjacencyMatrix) Inverse() AdjacencyMatrix {
	inverted := make(AdjacencyMatrix, len(graph))

	for from, toVector := range graph {
		for _, to := range toVector {
			inverted[to] = append(inverted[to], from)
		}
	}

	return inverted
}

// PathExistsDFS - проверяет существование пути в графе из вершины from в
// вершину to с помощью поиска в глубину.
func (graph AdjacencyMatrix) PathExistsDFS(from, to int) bool {
	visited := structs.NewBitSet(len(graph))
	stack := structs.Stack[int]{}
	stack.Push(from)

	for {
		v, ok := stack.Pop()
		if !ok {
			break
		}
		if v == to {
			return true
		}
		if !visited.IsSet(v) {
			visited.Set(v)
		}

		for i := len(graph[v]) - 1; i >= 0; i-- {
			w := graph[v][i]
			if !visited.IsSet(w) {
				stack.Push(w)
			}
		}
	}

	return false
}

// WalkDFS - обходит вершины графа начиная с вершины vertex с помощью поиска
// в глубину и вызывает для каждой callback-функцию.
func (graph AdjacencyMatrix) WalkDFS(vertex int, f func(vertex int)) {
	visited := structs.NewBitSet(len(graph))
	stack := structs.Stack[int]{}
	stack.Push(vertex)

	for {
		v, ok := stack.Pop()
		if !ok {
			break
		}
		if !visited.IsSet(v) {
			f(v)
			visited.Set(v)
		}

		for i := len(graph[v]) - 1; i >= 0; i-- {
			w := graph[v][i]
			if !visited.IsSet(w) {
				stack.Push(w)
			}
		}
	}
}

type componentsSeeker struct {
	graph    AdjacencyMatrix
	inverted AdjacencyMatrix

	stack      structs.Stack[int]
	visited    structs.BitSet
	components [][]int
}

func newComponentsSeeker(graph AdjacencyMatrix) *componentsSeeker {
	return &componentsSeeker{
		graph:      graph,
		inverted:   graph.Inverse(),
		visited:    structs.NewBitSet(len(graph)),
		components: make([][]int, 0),
	}
}

func (s *componentsSeeker) Find() [][]int {
	for _, vector := range s.inverted {
		for _, u := range vector {
			if !s.visited.IsSet(u) {
				s.dfs1(u)
			}
		}
	}

	componentIndex := 0
	s.visited = structs.NewBitSet(len(s.graph))
	for {
		u, ok := s.stack.Pop()
		if !ok {
			break
		}
		if !s.visited.IsSet(u) {
			s.components = append(s.components, []int{})
			s.dfs2(u, componentIndex)
			componentIndex++
		}
	}

	return s.components
}

func (s *componentsSeeker) dfs1(vertex int) {
	s.visited.Set(vertex)
	for _, u := range s.inverted[vertex] {
		if !s.visited.IsSet(u) {
			s.dfs1(u)
		}
	}
	s.stack.Push(vertex)
}

func (s *componentsSeeker) dfs2(vertex int, index int) {
	s.visited.Set(vertex)
	for _, u := range s.graph[vertex] {
		if !s.visited.IsSet(u) {
			s.dfs2(u, index)
		}
	}
	s.components[index] = append(s.components[index], vertex)
}
