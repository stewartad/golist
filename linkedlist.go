package golist

import "fmt"

type node struct {
	data int
	next *node
}

type list struct {
	head *node
	tail *node
}

func NewList() list {
	return list{head: nil, tail: nil}
}

func (l *list) AppendToTail(d int) {
	tmp := &node{data: d, next: nil}
	if l.head == nil {
		l.head = tmp
		l.tail = tmp
	} else {
		l.tail.next = tmp
		l.tail = tmp
	}
}

func (l *list) AppendToHead(d int) {
	tmp := &node{data: d, next: l.head}
	if l.head == nil {
		l.head = tmp
		l.tail = tmp
	} else {
		l.head = tmp
	}
}

func (l *list) PrintList() {
	p := l.head
	if p == nil {
		return
	}
	for p != l.tail {
		fmt.Printf("%d, ", p.data)
		p = p.next
	}
	fmt.Printf("%d\n", l.tail.data)
}
