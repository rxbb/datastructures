package stack

import (
	"log"
	"github.com/rxbb/datastructures/linkedlist"
)

type Stack struct{
	List *linkedlist.DoublyLinkedList
}

func NewStack() *Stack{
	stack := new(Stack)
	stack.List = linkedlist.NewDLL()
	return stack
}


func (stack *Stack) Push(item interface{}) {
	stack.List.AddLast(item)
}

func (stack *Stack) Size() int {
	return stack.List.Size
}

func (stack *Stack) Pop() interface{}{
	if stack.List.Size == 0 {
		log.Fatal("Empty List")
	}
	return stack.List.RemoveLast()
}

func (stack *Stack) Iterator() (func()bool,func()interface{}){
	return stack.List.Iterator()
}


func (stack *Stack) Peek() interface{} {
	return stack.List.PeekLast()
}
