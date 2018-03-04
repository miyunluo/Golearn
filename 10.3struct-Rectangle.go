// 定义一个 Rectangle 结构体，它的长和宽是 int 类型，并定义方法 Area() 和 Primeter()，然后进行测试
package main 
import "fmt"

type Rectangle struct {
	length, width int
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func (r *Rectangle) Perimeter() int {
	return 2*(r.length + r.width)
}

func main() {
	r := Rectangle{1,2}
	fmt.Println("Rectangle is: ", r)
	fmt.Println("Rectangle area is: ", r.Area())
	fmt.Println("Rectangle perimeter is: ", r.Perimeter())
}