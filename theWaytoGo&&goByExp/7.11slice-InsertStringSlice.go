// 将切片插入到另一个切片的指定位置
package main

import "fmt"

func main() {
	s1 := []string{"A", "A", "A", "A", "A", "A", "A"}
	s2 := []string{"B", "B", "B"}

	result0 := InsertSrtingSlice(s1, s2, 0)
	fmt.Println(result0)

	result1 := InsertSrtingSlice(s1, s2, 3)
	fmt.Println(result1)
}

func InsertSrtingSlice(s1, s2 []string, index int) []string {
	tmp := make([]string, len(s1)+len(s2))
	at := copy(tmp, s1[:index])
	at += copy(tmp[at:],s2)
	copy(tmp[at:],s1[index:])
	return tmp
}