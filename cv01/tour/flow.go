package main

import (
	"fmt"
	"runtime"
)

func main() {
	forLoop()
	switchStatement()
	deferStack()
}

func forLoop() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func switchStatement() {
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}
}

func deferStack() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
