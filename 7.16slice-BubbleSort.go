package main
import "fmt"

func main() {
	test := []int{2, 6, 4, -10, 8, 89, 12, 68, -45, 37}
	bubbleSort(test)
	fmt.Println("after sort:  ", test)
}

func bubbleSort(ls []int) {
	for pass := 1; pass < len(ls); pass ++ {
		for i := 0; i < len(ls)-pass; i++ {
			if ls[i] > ls[i+1] {
				ls[i], ls[i+1] = ls[i+1], ls[i]
			}
		}
	}
}