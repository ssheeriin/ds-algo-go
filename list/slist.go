package list

import (
	"fmt"
)

type slnode struct {
	data int
	next *slnode
}

type SLinkedList struct {
	head *slnode
	tail *slnode
}

func (ll *SLinkedList) AddToHead(data int) {
	ll.head = &slnode{data: data, next: ll.head}

	if ll.tail == nil {
		ll.tail = ll.head
	}
}

func (ll *SLinkedList) Append(data int) {
	if ll.tail == nil {
		ll.AddToHead(data)
	} else {
		newNode := &slnode{data: data, next: nil}
		ll.tail.next = newNode
		newNode = ll.tail
	}

}

func (ll *SLinkedList) Print() {
	curr := ll.head
	for curr != nil {
		fmt.Printf("%d %s", curr.data, " ")
		curr = curr.next
	}
	fmt.Println()
}
