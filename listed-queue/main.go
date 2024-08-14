package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Integer | constraints.Float
}

type Node[T Numeric] struct {
	data T
	next *Node[T]
}

type LinkedQueue[T Numeric] struct {
	head       *Node[T]
	tail       *Node[T]
	countNodes int
}

func New[T Numeric](node *Node[T]) *LinkedQueue[T] {
	return &LinkedQueue[T]{
		head:       node,
		tail:       node,
		countNodes: 1,
	}
}

func (q *LinkedQueue[T]) get() {
	if q.head == nil {
		fmt.Println("error")
		return
	}
	current := q.head
	q.head = current.next
	q.countNodes--
	fmt.Println(current.data)
}

func (q *LinkedQueue[T]) PushBack(x T) {
	node := &Node[T]{data: x, next: nil}
	q.tail.next = node
	q.tail = node
	q.countNodes++
}

func (q *LinkedQueue[T]) Size() {
	fmt.Println(q.countNodes)
}

func main() {
	node := Node[int]{data: 100, next: nil}
	queue := New(&node)
	queue.Size()
	queue.PushBack(123)
	queue.PushBack(12)
	queue.PushBack(13)
	queue.PushBack(23)
	queue.PushBack(1)
	queue.PushBack(3)
	queue.get()
	queue.get()
	queue.Size()
	fmt.Println(queue.head.data)
	fmt.Println(queue.tail.data)
}
