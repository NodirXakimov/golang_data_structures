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

func (l LinkedList) addToList(n *node) {
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
	var a int
	fmt.Println(a)
}
