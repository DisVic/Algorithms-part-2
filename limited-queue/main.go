package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/constraints"
)

type Numberic interface {
	constraints.Float | constraints.Integer
}

type MyQueueSized[T Numberic] struct {
	len   int
	queue []T
}

func New[T Numberic](len int) *MyQueueSized[T] {
	return &MyQueueSized[T]{
		len: len,
	}
}

func (q *MyQueueSized[T]) push(x T) {
	if q.len > len(q.queue) {
		q.queue = append(q.queue, x)
	} else {
		fmt.Println("Очередь переполнена!")
		return
	}
}

func (q *MyQueueSized[T]) pop() {
	if len(q.queue) == 0 {
		fmt.Println("Очередь пуста")
		return
	}
	number := q.queue[0]
	q.queue = q.queue[1:]
	fmt.Println("Удалённое число: ", number)
}

func (q *MyQueueSized[T]) peek() {
	if len(q.queue) == 0 {
		fmt.Println("Очередь пуста")
		return
	}
	fmt.Println("Первое число в очереди: ", q.queue[0])
}

func (q *MyQueueSized[T]) size() int {
	fmt.Println("Размер текущей очереди: ", len(q.queue))
	fmt.Println("Допустимый размер очереди: ", q.len)
	return len(q.queue)
}

func main() {
	queue := New[int](10)
	t := time.Now()
	queue.size()
	queue.push(55)
	queue.push(357)
	queue.push(7)
	queue.push(9)
	queue.push(91)
	queue.push(789)
	queue.push(71)
	queue.push(1000)
	queue.push(645)
	queue.push(497)
	queue.push(1)
	queue.push(778)
	queue.peek()
	queue.pop()
	queue.size()
	fmt.Println(time.Since(t))
	fmt.Println(queue.queue)
}
