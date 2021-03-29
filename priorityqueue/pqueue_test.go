package priorityqueue

import (
	"fmt"
	"testing"
)


func TestPQ(t *testing.T){
	pq := new(PQueue)
	Init(pq)
	Push(pq,3)
	Push(pq,7)
	Push(pq,8)
	Push(pq,10)
	Push(pq,1)
	Push(pq,2)
	Push(pq,32)
	Push(pq,11)
	first := Pop(pq)
	fmt.Println(pq,first)
}
