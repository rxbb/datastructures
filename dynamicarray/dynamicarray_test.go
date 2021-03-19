package dynamicarray

import "testing"


func TestDA(t *testing.T){
	arr := NewIntArray()
	arr.Push(1)
	arr.Push(3)
	arr.Push(5)
	arr.Push(11)
	arr.Push(112)
	if arr.Length() != 5 {
		t.Errorf("Array's Length was Incorrect : {%d} --> {%d} ",arr.Length(),5)
	}

	lastItem := arr.Pop() 
	if lastItem!= 112 || arr.Length() != 4 {
		t.Errorf("Last item or Array's Length was Incorrect : Last item was {%d}  and Lenght was {%d}",lastItem,arr.Length())
	}
	removedItem := arr.RemoveAt(2)
	if removedItem != 5 || arr.Length() != 3 {
		t.Errorf("Removed item was Incorrect or Array's Length was incorrect : {%d} --> {%d} | Array Length was {%d} Must been {%d}",removedItem,5,arr.Length(),3)
	}
}
