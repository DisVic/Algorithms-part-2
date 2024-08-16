package main

import (
	"fmt"
)

type Deque[T any] struct {
	deq []T
}

func (d *Deque[T]) PushFront(item T) {
	d.deq = append([]T{item}, d.deq...)
}

func (d *Deque[T]) PushBack(item T) {
	d.deq = append(d.deq, item)
}

func (d *Deque[T]) PopBack() {
	if len(d.deq) == 0 {
		fmt.Println("Двойная очередь пуста!")
		return
	}
	fmt.Println(d.deq[len(d.deq)-1])
	d.deq = d.deq[:len(d.deq)-1]
}

func (d *Deque[T]) PopFront() {
	if len(d.deq) == 0 {
		fmt.Println("Двойная очередь пуста!")
		return
	}
	fmt.Println(d.deq[0])
	d.deq = d.deq[1:]
}

func main() {
	deq := new(Deque[string])
	deq.PushFront("1")
	deq.PushFront("2")
	deq.PushFront("3")
	deq.PushFront("4")
	deq.PushFront("5")
	deq.PushFront("6")
	deq.PopBack()
	deq.PopFront()
	fmt.Println(deq.deq)
}
