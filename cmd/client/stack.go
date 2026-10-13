package client

type Stack[T any] struct {
	values []T
	limit  int
}

func (s *Stack[T]) Pop() {
	s.values = s.values[1:]
}

func (s *Stack[T]) Push(val T) {
	if len(s.values)+1 > s.limit {
		s.Pop()
	}
	s.values = append(s.values, val)
}

func (s *Stack[T]) Each(f func(item T)) {
	for i := len(s.values) - 1; i >= 0; i-- {
		f(s.values[i])
	}
}
func (s *Stack[T]) EachBackwards(f func(item T)) {
	for _, v := range s.values {
		f(v)
	}
}
