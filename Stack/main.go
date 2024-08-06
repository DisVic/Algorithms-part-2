package main

import (
	"fmt"
)

type Stack[T any] struct {
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
	var max T
	switch any(s.stack).(type) {
	case T == string:
		for i := 0; i < len(s.stack)-1; i++ {
			if len(s.stack[i]) < len(s.stack[i+1]) {

			}
		}
	}

}

func main() {
	stk := new(Stack[string])
	stk.push("lol")
	stk.push("pop")
	stk.pop()

	fmt.Printf("%v ", stk.stack)
}
