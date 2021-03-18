package dynamic_array

type IntArray struct {
	Arr []int
	N   int
}

func NewIntArray() *IntArray {
	intArr := new(IntArray)
	intArr.Arr = make([]int, 1)
	intArr.N = 0
	return intArr
}


func (a *IntArray) IsEmpty() bool{
	return a.N == 0
}

func (ia *IntArray) Length() int {
	return len(ia.Arr)
	// or ia.N + 1
}

func (ia *IntArray) resize(cp int) {
	newArr := make([]int, cp)
	for i, v := range ia.Arr {
		newArr[i] = v
	}
	ia.Arr = newArr
}

func (ia *IntArray) Push(item int) {
	if ia.N == len(ia.Arr) {
		ia.resize(2 * len(ia.Arr))
	}
	ia.Arr[ia.N] = item
	ia.N++
}

func (ia *IntArray) Pop() int {
	ia.N--
	item := ia.Arr[ia.N]
	ia.Arr[ia.N] = 0
	return item
}
