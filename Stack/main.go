package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Float | constraints.Integer
}

type Stack[T Numeric] struct {
	max   T
	stack []T
}

func (s *Stack[T]) push(input T) {
	if input > s.max {
		s.max = input
	}
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
	return s.max
}

func main() {
	stk := new(Stack[int])
	stk.push(10)
	stk.push(5)
	fmt.Println(stk.get_max())
	stk.pop()

	fmt.Printf("%v ", stk.stack)
}
