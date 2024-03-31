package calculate

func PerformOperation(a int, b int) (addNum int, sub int, mul int, div int) {
	const val float64 = 0.00
	addNum = Calculate(Add, a, b)
	sub = Calculate(Subtract, a, b)
	mul = Calculate(Multiply, a, b)
	div = Calculate(Divide, a, b)
	return addNum, sub, mul, div
}

func Calculate(callBack func(a int, b int) (c int), num1 int, num2 int) (output int) {
	output = callBack(num1, num2)
	return output
}

func Add(a int, b int) (c int) {

	// d := a+ b
	c = a + b
	return c
}
func Subtract(a int, b int) (c int) {

	c = a - b
	return c
}
func Multiply(a int, b int) (c int) {

	c = a * b
	return c
}
func Divide(a int, b int) (c int) {

	c = a / b
	return c
}
