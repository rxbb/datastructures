package dynamicarray

import "fmt"

type Array struct {
	Arr []interface{}
	N   int
}

func NewArray() *Array {
	intArr := new(Array)
	intArr.Arr = make([]interface{}, 1)
	intArr.N = 0
	return intArr
}


func (a *Array) IsEmpty() bool{
	return a.N == 0
}

func (a *Array) Length() int {
	return a.N
}

func (a *Array) resize(cp int) {
	newArr := make([]interface{}, cp)
	for i, v := range a.Arr {
		newArr[i] = v
	}
	a.Arr = newArr
}

func (a *Array) Push(item interface{}) {
	if a.N == len(a.Arr) {
		a.resize(2 * len(a.Arr))
	}
	a.Arr[a.N] = item
	a.N++
}

func (a *Array) Pop() interface{} {
	a.N--
	item := a.Arr[a.N]
	a.Arr[a.N] = 0
	return item
}

func (a *Array) RemoveAt(index int) interface{} {
	data := a.Arr[index]
	newArr := make([]interface{},a.N - 1)
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
