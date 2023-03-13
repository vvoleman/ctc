package main

import (
	"fmt"
	"math"
)

type Vert struct {
	X, Y float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

func (v Vert) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Person struct {
	Name string
	Age  string
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	v := Vert{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	a := Person{"Arthur Dent", "42"}
	z := Person{"Zaphod Beeblebrox", "9001"}
	fmt.Println(a, z)
}

func errors() {

}
