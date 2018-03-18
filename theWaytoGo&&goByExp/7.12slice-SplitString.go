// 原始字符串 str 和分割索引 i，然后返回两个分割后的字符串
package main

import "fmt"

func main() {
	str := "golang"
	a,b := Split(str, i)
	fmy.Printf("String %s split at position %d is %s / %s\n", str, i, a, b)
}

func Split(s string, pos int) (string, string) {
	return s[0:pos], s[pos:]
}