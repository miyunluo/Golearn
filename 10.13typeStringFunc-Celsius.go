package main 

import (
	"fmt"
	"strconv"
)

type Celsius float64

func (c Celsius) String() string {
	return "The temprature is: " + strconv.FormatFloat(float64(c), 'f', 1, 32) + " °C"
}

func main() {
	var c Celsius = 10.111
	fmt.Println(c)
}