package main

import (
	"fmt"
	"test/calculate"
)

func init() {
	
	fmt.Println("____________________________________________")
}
func main() {
	add, sub, mul, div := calculate.PerformOperation(4, 2)
	fmt.Println("addition", add)
	fmt.Println("Subtraction", sub)
	fmt.Println("Multiplication", mul)
	fmt.Println("Division", div)
	fmt.Println("PublicVariable",calculate.PublicVariable)
	calculate.Looping()
	calculate.SwitchStatement()
	calculate.DeferImplementation()
	calculate.Structure()
	calculate.ArrayAndSlice()


}
