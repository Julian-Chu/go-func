package main

import "fmt"

func main() {
	c := &Calculator{}
	c.Calc(Add, 2)
	fmt.Println(c.Val)
}

type Calculator struct {
	Val float64
}

func (c *Calculator) Calc(op Operator, b float64) {
	c.Val = op(c.Val, b)
}

type Operator func(a, b float64) float64

var (
	Add Operator = func(a, b float64) float64 {
		return a + b
	}
)
