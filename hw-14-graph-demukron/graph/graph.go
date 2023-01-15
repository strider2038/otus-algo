package graph

import (
	"errors"
)

var ErrNotSortable = errors.New("graph is not sortable")

type Vector []int

// AdjacencyArray - представление графа в виде массива векторов смежности.
type AdjacencyArray []Vector

// SortByDemukron - топологическая сортировка алгоритмом Демукрона.
// Возвращает ошибку ErrNotSortable в случае если в графе обнаружены циклы и он
// не может быть отсортирован.
func (graph AdjacencyArray) SortByDemukron() ([][]int, error) {
	matrix := graph.AdjacencyMatrix()
	incomingDegrees := matrix.IncomingDegrees()

	sorted := make([][]int, 0)

	for incomingDegrees.Valid() {
		vertices := incomingDegrees.ZeroVertices()
		if len(vertices) == 0 {
			return nil, ErrNotSortable
		}

		sorted = append(sorted, vertices)
		for _, vertex := range vertices {
			incomingDegrees[vertex] = -1
			incomingDegrees = incomingDegrees.Sub(matrix[vertex])
		}
	}

	return sorted, nil
}

// AdjacencyMatrix - преобразует представление графа в матрицу смежности.
func (graph AdjacencyArray) AdjacencyMatrix() AdjacencyMatrix {
	matrix := make(AdjacencyMatrix, len(graph))

	for from := 0; from < len(graph); from++ {
		matrix[from] = make([]int, len(graph))
		for _, to := range graph[from] {
			matrix[from][to] = 1
		}
	}

	return matrix
}

// AdjacencyMatrix - представление графа в виде матрицы смежности.
type AdjacencyMatrix [][]int

// IncomingDegrees - возвращает массив степеней входящих вершин.
func (matrix AdjacencyMatrix) IncomingDegrees() DegreeVector {
	incomingDegrees := make(DegreeVector, len(matrix))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			incomingDegrees[i] += matrix[j][i]
		}
	}

	return incomingDegrees
}

// DegreeVector - массив степеней вершин.
type DegreeVector []int

func (vector DegreeVector) Valid() bool {
	for _, degree := range vector {
		if degree >= 0 {
			return true
		}
	}

	return false
}

// ZeroVertices возвращает вершины с нулевой степенью.
func (vector DegreeVector) ZeroVertices() []int {
	vertices := make([]int, 0)

	for vertex, degree := range vector {
		if degree == 0 {
			vertices = append(vertices, vertex)
		}
	}

	return vertices
}

// Sub вычитает из текущего массива указанный и возвращает новый массив.
func (vector DegreeVector) Sub(subtrahend DegreeVector) DegreeVector {
	result := make(DegreeVector, len(vector))

	for i := 0; i < len(vector); i++ {
		result[i] = vector[i] - subtrahend[i]
	}

	return result
}
