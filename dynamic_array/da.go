package dynamic_array

import "fmt"

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

func (a *IntArray) Length() int {
	return a.N
}

func (a *IntArray) resize(cp int) {
	newArr := make([]int, cp)
	for i, v := range a.Arr {
		newArr[i] = v
	}
	a.Arr = newArr
}

func (a *IntArray) Push(item int) {
	if a.N == len(a.Arr) {
		a.resize(2 * len(a.Arr))
	}
	a.Arr[a.N] = item
	a.N++
}

func (a *IntArray) Pop() int {
	a.N--
	item := a.Arr[a.N]
	a.Arr[a.N] = 0
	return item
}

func (a *IntArray) RemoveAt(index int) int {
	data := a.Arr[index]
	newArr := make([]int,a.N - 1)
	i := 0
	j := 0
	fmt.Println(a.N)
	for i < a.N {
		if i == index{
			j--
		}else{
			newArr[j] = a.Arr[i]
		}
		i++
		j++
	}
	a.Arr = newArr
	a.N--
	return data
}
