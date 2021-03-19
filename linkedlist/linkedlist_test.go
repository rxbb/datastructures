package linkedlist

import (
	"fmt"
	"testing"
)


func TestLinkedList(t *testing.T){
	list := NewDLL()
	list.AddLast("first item")
	list.AddLast("second item")
	list.AddLast("third item")
	list.AddLast("fourth item")
	if list.Size != 4 {
		t.Errorf("Size of List was incorrect --> %d",list.Size)
	}
	if list.Tail.Data != "fourth item"{
		t.Errorf("Tail of the List was incorrect --> %s",list.Tail.Data)
	}
	list.RemoveFirst()
	if list.Head.Data == "first item"{
		t.Errorf("Head of the List was incorrect after removing first node --> %s",list.Head.Data)
	}
	list.RemoveLast()
	if list.Tail.Data == "fourth item"{
		t.Errorf("Tail of the List was incorrect after removing last node --> %s",list.Tail.Data)
	}
	list.RemoveAt(1)
	if list.Head.Data != list.Tail.Data {
		t.Errorf("tail and head data must be equal after removing all nodes but one --> %s",list.Tail.Data)
	}
}


func TestIterator(t *testing.T){
	list := NewDLL()
	list.AddLast("first item")
	list.AddLast("second item")
	list.AddLast("third item")
	list.AddLast("fourth item")
	hasNext , next:= list.Iterator()
	for hasNext() {
		fmt.Println(next())
	}
}
