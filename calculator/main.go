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

type Operator string

var (
	Add Operator = "Add"
	Sub Operator = "Sub"
)

func (c *Calculator) Calc(op Operator, v float64) {
	switch op {
	case "Add":
		c.Val += v
		break
	case "Sub":
		c.Val -= v
		break
	}
}
