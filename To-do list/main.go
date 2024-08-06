package main

import (
	"fmt"
)

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) add(s string) {
	newNode := &Node{data: s, next: nil}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (l *LinkedList) remove(index int) {
	if index <= 0 {
		fmt.Println("неверный индекс элемента списка")
		return
	}
	current := l.head
	if index == 1 {
		l.head = current.next
		return
	}
	if current == nil || current.next == nil {
		fmt.Println("Длина связного списка <2")
		return
	}
	var prevNode *Node
	for i := 1; i != index; i++ {
		if current.next != nil {
			prevNode = current
			current = current.next
		} else {
			fmt.Println("нет такого элемента в списке")
			return
		}
	}
	prevNode.next = current.next
}

func (l *LinkedList) getElement(index int) {
	if index <= 0 {
		fmt.Println("неверный индекс элемента списка")
		return
	}
	current := l.head
	if current == nil {
		fmt.Println("В списке нет значений")
		return
	}
	for i := 1; i != index; i++ {
		if current.next != nil {
			current = current.next
		} else {
			fmt.Println("нет такого элемента в списке")
			return
		}
	}
	fmt.Println("Вот ваш элемент: " + current.data)
}

func (l LinkedList) printLinkedList() {
	current := l.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func main() {
	list := new(LinkedList)
	list.add("first")
	list.add("second")
	list.add("third")
	list.printLinkedList()
	list.getElement(3)
	list.remove(2)
	list.printLinkedList()
	list.getElement(1)
}
