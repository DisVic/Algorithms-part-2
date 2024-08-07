package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Stack[T constraints.Ordered] struct {
	stack []T
}

func (s *Stack[T]) push(input T) {
	s.stack = append(s.stack, input)
}

func (s *Stack[T]) pop() {
	if len(s.stack) == 0 {
		fmt.Println("стэк пуст")
		return
	}
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *Stack[T]) get_max() T {
	if len(s.stack) == 0 {
		fmt.Println("стэк пуст")
		var zero T
		return zero
	}
	max := s.stack[0]
	for _, v := range s.stack {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	stk := new(Stack[int])
	stk.push(10)
	stk.push(5)
	fmt.Println(stk.get_max())
	stk.pop()

	fmt.Printf("%v ", stk.stack)
}
