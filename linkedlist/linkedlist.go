package linkedlist

import (
	"log"
)

type Node struct {
	Data interface{}
	Next *Node
	Prev *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func NewNode(data interface{}, next *Node, prev *Node) *Node {
	n := new(Node)
	n.Next = next
	n.Prev = prev
	n.Data = data
	return n
}

func NewDLL() *DoublyLinkedList {
	dll := new(DoublyLinkedList)
	dll.Head = nil
	dll.Tail = nil
	dll.Size = 0
	return dll
}

func (dll *DoublyLinkedList) IsEmpty() bool {
	return dll.Size == 0
}

func (dll *DoublyLinkedList) Clear() {
	trav := dll.Head
	for trav != nil {
		next := trav.Next
		trav.Prev = nil
		trav.Next = nil
		trav = next
	}
	dll.Size = 0
	dll.Head = nil
	dll.Tail = nil
	trav = nil
}

func (dll *DoublyLinkedList) AddLast(data interface{}) {
	if dll.IsEmpty() {
		dll.Head = NewNode(data, nil, nil)
		dll.Tail = dll.Head
	} else {
		dll.Tail.Next = NewNode(data, nil, dll.Tail)
		dll.Tail = dll.Tail.Next
	}
	dll.Size++
}

func (dll *DoublyLinkedList) AddFirst(data interface{}) {
	if dll.IsEmpty() {
		dll.Head = NewNode(data, nil, nil)
		dll.Tail = dll.Head
	} else {
		dll.Head.Prev = NewNode(data, dll.Head, nil)
		dll.Head = dll.Head.Prev
	}
	dll.Size++
}

func (dll *DoublyLinkedList) PeekLast() interface{} {
	if dll.IsEmpty() {
		log.Fatal("List is Empty")
	}
	return dll.Tail.Data
}

func (dll *DoublyLinkedList) PeekFirst() interface{} {
	if dll.IsEmpty() {
		log.Fatal("List is Empty")
	}
	return dll.Head.Data
}

func (dll *DoublyLinkedList) RemoveFirst() interface{} {
	if dll.IsEmpty() {
		log.Fatal("List is Empty")
	}
	data := dll.Head.Data
	dll.Head = dll.Head.Next
	dll.Size--
	if dll.IsEmpty() {
		dll.Tail = nil
	} else {
		dll.Head.Prev = nil
	}
	return data
}

func (dll *DoublyLinkedList) RemoveLast() interface{} {
	if dll.IsEmpty() {
		log.Fatal("List is Empty")
	}
	data := dll.Tail.Data
	dll.Tail = dll.Tail.Prev
	dll.Size--
	if dll.IsEmpty() {
		dll.Head = nil
	} else {
		dll.Tail.Next = nil
	}
	return data
}

func (dll *DoublyLinkedList) Remove(node *Node) interface{} {
	if node.Prev == nil {
		return dll.RemoveFirst()
	}
	if node.Next == nil {
		return dll.RemoveLast()
	}

	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
	// Get node data to return it
	data := node.Data
	// deallocate the memory
	node.Data = nil
	node.Next = nil
	node.Prev = nil
	dll.Size--
	return data
}

func (dll *DoublyLinkedList) RemoveAt(index int) interface{} {
	if index < 0 || index >= dll.Size {
		log.Fatal("Illegal index")
	}
	var i int
	var trav *Node
	if index < dll.Size/2 {
		i = 0
		trav = dll.Head
		for i != index {
			trav = trav.Next
			i++
		}
	} else {
		i = dll.Size - 1
		trav = dll.Tail
		for i != index {
			trav = trav.Prev
			i--
		}
	}
	return dll.Remove(trav)
}

func (dll *DoublyLinkedList) Iterator() (func() bool, func() interface{}) {
	trav := dll.Head
	hasNext := func() bool {
		return trav != nil
	}
	next := func() interface{} {
		data := trav.Data
		trav = trav.Next
		return data
	}
	return hasNext, next
}
