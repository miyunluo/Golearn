package main
import (
	"fmt"
	"unsafe"
)

func main() {
	var i int = 10
	size := unsafe.Sizeof(i)
	fmt.Println("Size of int is: ", size)
}

