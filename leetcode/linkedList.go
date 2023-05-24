package main

import "fmt"

type Node struct {
	information interface{}
	next        *Node
}

type List struct {
	head *Node
}

func (l *List) Insert(d interface{}) {
	list := &Node{information: d, next: nil}
	if l.head == nil {
		l.head = list
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		p.next = list
	}
}

func Show(l *List) {
	p := l.head
	for p != nil {
		fmt.Printf("-> %v", p.information)
		p = p.next
	}
}

func main() {
	sl := List{}
	for i := 0; i < 10; i++ {
		sl.Insert(i)
	}
	Show(&sl)
	fmt.Println()
}
