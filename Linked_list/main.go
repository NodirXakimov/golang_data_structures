package main

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	main *node
	tail *node
	size int
}

func (l *LinkedList) addToList(n *node) {
	if l.main == nil {
		l.main = n
	}
	if l.tail != nil {
		l.tail.next = n
	}
	l.tail = n
	l.size++
	fmt.Println("Test", l.size, l.tail.data)
	l.tail.next = nil
}

func (l *LinkedList) pushBack(n *node) {
	l.tail.next = n
	l.tail = n
	l.size++
}

func main() {
	fmt.Println("Hello Linked List")
	list := LinkedList{}
	for i := 1; i < 10; i++ {
		n := node{
			data: i,
		}
		list.addToList(&n)
	}

	fmt.Println(list.size)
	fmt.Println(list.main.data)
	fmt.Println(list.tail.data)

	n := node{
		data: 100,
	}
	list.pushBack(&n)

	fmt.Println(list.size)
	fmt.Println(list.main.data)
	fmt.Println(list.tail.data)
	fmt.Println(list.tail.next)
}
