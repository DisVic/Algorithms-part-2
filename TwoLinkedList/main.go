package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack("first")
	l.PushBack("second")
	l.PushBack("third")
	l.PushBack("fourth")
	l.PushBack("fifth")
	l = reverseLinkedList(l)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
}

func reverseLinkedList(l *list.List) *list.List {
	newList := list.New()
	for e := l.Back(); e != nil; e = e.Prev() {
		newList.PushBack(e.Value)
	}
	return newList
}
