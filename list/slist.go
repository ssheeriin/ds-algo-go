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
}

func (ll *SLinkedList) AddToHead(data int) {
	ll.head = &slnode{data: data, next: ll.head}
}

func (ll *SLinkedList) Print() {
	curr := ll.head
	for curr != nil {
		fmt.Printf("%d %s", curr.data, " ")
		curr = curr.next
	}
	fmt.Println()
}
