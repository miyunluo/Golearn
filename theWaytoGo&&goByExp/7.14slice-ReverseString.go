// 反转字符串，即将 “Google” 转换成 “elgooG”（提示：使用 []byte 类型的切片)
// 如果您使用两个切片来实现反转，请再尝试使用一个切片（提示：使用交换法）
package main

import "fmt"

func main() {
	str := "Google"
	fmt.Printf("The reversed string is %s \n", reverse(str))
}

func reverse(s string) string {
	tmp := []byte(s)
	j := 0
	for i := len(s)-1; i >= 0; i-- {
		tmp[j] = s[i]
		j++
	}
	return string(tmp)
}