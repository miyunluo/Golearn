package main 

import (
	"fmt"
	"container/list"
)

func main() {
	ls := list.New()
	ls.PushBack(100)
	ls.PushBack(101)
	ls.PushBack(102)

	for e := ls.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}