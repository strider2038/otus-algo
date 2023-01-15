package graph

type Stack[T any] struct {
	top *StackItem[T]
}

type StackItem[T any] struct {
	value T
	next  *StackItem[T]
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack[T]) Push(value T) {
	s.top = &StackItem[T]{value: value, next: s.top}
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	value := s.top.value
	s.top = s.top.next

	return value, true
}
