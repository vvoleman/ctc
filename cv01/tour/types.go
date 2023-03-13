package main

import "fmt"

func main() {
	//pointers()
	//structs()
	//arrays()
	//ranges()
	maps()
}

func pointers() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

type Vertex struct {
	X int
	Y int
}

func structs() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	p := &v
	p.X = 1e9
	fmt.Println(p.X)
}

func arrays() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	fmt.Printf("Lenght: %d, cap: %d\n", len(a), cap(a))

	a = append(a, "ahoj")
	fmt.Println(a)
}

func ranges() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

type Position struct {
	Lat, Long float64
}

func maps() {
	m := make(map[string]Position)
	m["Bell"] = Position{40.684, -74.399}
	fmt.Println(m["Bell"])

}
