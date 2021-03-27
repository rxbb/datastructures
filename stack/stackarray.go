package stack

import (
	"github.com/rxbb/datastructures/dynamicarray"
)

type StackArr struct {
	List *dynamicarray.Array
}

func NewStackArr() *StackArr {
	stack := new(StackArr)
	stack.List = dynamicarray.NewArray()
	return stack
}

func (stack *StackArr) Size() int {
	return stack.List.Length()
}

func (stack *StackArr) Push(item int) {
	stack.List.Push(item)
}

func (stack *StackArr) Pop() interface{} {
	return stack.List.Pop()
}
