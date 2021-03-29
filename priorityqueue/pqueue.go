package priorityqueue


type Heap interface {
	Swap(i, j int)
	Len() int
	Less(i, j int) bool
	Pop() interface{}
	Push(x interface{})
}

type PQueue []int

func (pq *PQueue) Len() int {
	return len(*pq)
}

func (pq *PQueue) Push(x interface{}) {
	*pq = append(*pq,x.(int))
}

func (pq *PQueue) Pop()interface{}{
	old := *pq
	item := old[len(old)-1]
	old[len(old)-1] = 0
	*pq = old[0 : len(old)-1]
	return item
}

func (pq *PQueue) Less(i, j int) bool {
	// we want the lowest to be root of the heap
	// so we use less tha operator here.
	return (*pq)[i] < (*pq)[j]
}

func (pq *PQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func Init(h Heap) {
	n := h.Len() - 1
	for i := n/2 - 1; i >= 0; i-- {
		Down(h, i, n)
	}
}
func Push(h Heap, x int) {
	h.Push(x)
	Up(h, h.Len()-1)
}
func Pop(h Heap) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	Down(h, 0, n)
	return h.Pop()
}
func Down(h Heap, x int, n int) {
	var j int
	for {
		j1 := 2*x + 1
		if j1 >= n || j1<0 {
			break
		}
		j = j1
		if j2:=j1+1; j2<n && h.Less(j2,j1) {
			j = j2
		}
		h.Swap(x,j)
		x = j
	}
}

func Up(h Heap, x int) {
	for {
		p := (x - 1) / 2
		if p == x || !h.Less(x, p) {
			break
		}
		h.Swap(x, p)
		x = p
	}
}
