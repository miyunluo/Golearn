// 将从 start 到 end 索引的元素从切片 中移除
package main

import "fmt"

func main(){
	s := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	result := RemoveStringSlice(s, 2, 4)
	fmt.Println(result)
}

func RemoveStringSlice(s []string, begin, end int) []string{
	tmp := make([]string, len(s)-(end - begin))
	at := copy(tmp, s[:begin])
	copy(tmp[at:], s[end:])
	return tmp
}
