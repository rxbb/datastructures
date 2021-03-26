package queue

import "testing"


func TestQueue(t *testing.T){
	q := NewQueue()
	q.Enqueue("first")
	q.Enqueue("second")
	q.Enqueue("third")
	if q.IsEmpty(){
		t.Error("Queue is Not Empty")
	}
	if first:=q.Dequeue(); first != "first"{
		t.Errorf("first item was'nt correnct")
	}
	if peek:=q.Peek();peek!="second"{
		t.Errorf("peek item was'nt correnct")
	}
}
