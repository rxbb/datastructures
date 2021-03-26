package queue

import "github.com/rxbb/datastructures/linkedlist"

type Queue struct {
	List *linkedlist.DoublyLinkedList
}

func NewQueue() *Queue{
	res := new(Queue)
	res.List = linkedlist.NewDLL()
	return res
}

func (q *Queue) Enqueue(item interface{}) {
	q.List.AddLast(item)
}

func (q *Queue) Dequeue() interface{}{
	return q.List.RemoveFirst()
}

func (q *Queue) IsEmpty() bool{
	return q.List.IsEmpty()
}

func (q *Queue) Peek() interface{}{
	return q.List.PeekFirst()
}


func (q *Queue) Iterator() (func()bool , func()interface{}){
	return q.List.Iterator()
}
