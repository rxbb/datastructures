package stack

import "testing"

func TestStack(t *testing.T){
	stack := NewStack()
	stack.Push("first")
	stack.Push("second")
	if stack.Peek() != "second" {
		t.Errorf("incorrect peek --> %s",stack.Peek())
	}
	last := stack.Pop()
	if last != "second" {
		t.Errorf("incorrect last item --> %s",last)
	}
	if stack.Peek() != "first" {
		t.Errorf("incorrect peek after removing last item --> %s",stack.Peek())
	}
	if stack.Size() != 1{
		t.Errorf("incorrect Size--> %d",stack.Size())
	}
}
