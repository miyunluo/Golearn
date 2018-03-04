package main 

import(
	"./float64"
	"fmt"
)

func main() {
	f := float64.NewFloat64Array()
	f.Fill(10)
	fmt.Println("Initial %s\n", f)
	float64.Sort(f)
	fmt.Printf("After sorting %s\n", f)
}