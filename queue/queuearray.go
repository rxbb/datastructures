package queue

const INIT_CAP = 8

type array struct {
	list  []interface{}
	n     int
	front int
	back  int
}

type QueueArray struct {
	List *array
}

func newArray() *array {
	ar := new(array)
	ar.n = 0
	ar.back = 0
	ar.front = 0
	ar.list = make([]interface{}, INIT_CAP)
	return ar
}

func NewQueueArray() *QueueArray {
	q := new(QueueArray)
	q.List = newArray()
	return q
}

func (arr *array) resize(capi int) {
	newArr := make([]interface{}, arr.n+1)
	for i := 0; i < arr.n; i++ {
		newArr[i] = arr.list[(i+arr.front)%len(arr.list)]
	}
	arr.list = newArr
	arr.front = 0
	arr.back = arr.n
}

func (arr *array) enqueue(item interface{}) {
	if arr.n == len(arr.list) {
		arr.resize(len(arr.list))
	}
	arr.list[arr.back] = item
	arr.back++
	if arr.back == len(arr.list) {
		arr.back = 0
	}
	arr.n++
}

func (arr *array) dequeue() interface{} {
	item := arr.list[arr.front]
	arr.list[arr.front] = nil
	arr.n--
	arr.front++
	if arr.front == len(arr.list) {
		arr.front = 0
	}
	if arr.n > 0 && arr.n == len(arr.list)/4 {
		arr.resize(len(arr.list) / 2)
	}
	return item
}

func (q *QueueArray) Enqueue(item interface{}) {
	q.List.enqueue(item)
}

func (q *QueueArray) Dequeue() interface{} {
	return q.List.dequeue()
}
