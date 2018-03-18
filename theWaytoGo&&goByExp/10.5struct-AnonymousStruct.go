// 匿名字段
package main 

import "fmt"

type S struct {
	x float32
	int 
	string
}

func main() {
	s := S{1.11, 0, "string"}
	fmt.Println(s.x, s.int, s.string)
	fmt.Println(s)
}