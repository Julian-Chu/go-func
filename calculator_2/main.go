package main

import "fmt"

func main() {
	c := &Calculator{}
	c.Calc(Add{}, 2)
	fmt.Println(c.Val)
}

type Calculator struct {
	Val float64
}

func (c *Calculator) Calc(operator OperatorFunc, val float64) {
	c.Val = operator.Do(c.Val, val)
}

type OperatorFunc interface {
	Do(a, b float64) (res float64)
}

type Add struct {
}

func (op Add) Do(a, b float64) (res float64) {
	return a + b
}
