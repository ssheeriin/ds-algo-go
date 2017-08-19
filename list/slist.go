package list

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
