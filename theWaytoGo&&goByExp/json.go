package main 

import(
	"fmt"
	"os"
	"encoding/json"
)

type response1 struct{
	Page int
	Fruits []string
}
type response2 struct{
	Page int
	Fruits []string
}

// json.Marshal() 将对象转换为JSON 
// func Marshal(v  interface{}) ([]byte, error)

// json.Unmarshal() 反序列化
//func Unmarshal(data []byte, v interface{}) error

// 解码任意数据
// json 包使用 map[string]interface{} 和 []interface{} 储存任意的 JSON 对象和数组

func main(){
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))

    slcD := []string{"apple", "peach", "pear"}
    slcB,_ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))
}