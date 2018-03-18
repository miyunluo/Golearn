package main

import "fmt"

func main() {
	ls := []int{0, 1, 2, 3, 4, 5, 6, 7}
	mf := func(i int) int {
		return i*10
	}
	fmt.Printf("%v\n", mapFunc(mf, ls))
}

func mapFunc(mf func(int) int, ls []int) []int {
	result := make([]int, len(ls))
	for i, v := range ls{
		result[i] = mf(v)
	}
	return result
}