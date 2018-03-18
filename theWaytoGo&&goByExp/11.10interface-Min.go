package main 

import "fmt"

type Miner interface {
	Len() int
	Less(i,j int) bool
	ValIx(ix int) interface{}
}

func Min(data Miner) interface{} {
	min := data.ValIx(0)
	for i := 1; i< data.Len(); i++ {
		if data.Less(i, i-1) {
			min = data.ValIx(i)
		}
	}
	return min
}

type IntArray []int
func (Ia IntArray) Len() int {return len(Ia)}
func (Ia IntArray) Less(i,j int) bool {return Ia[i]<Ia[j]}
func (Ia IntArray) ValIx(ix int) interface{} {return Ia[ix]}

func main() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0}
	intarray := IntArray(data)
	min := Min(intarray)
	fmt.Printf("The minimum of the array is: %v\n", min)
}