package main

import "fmt"

const (
	ABS = iota
	B
	C
)

var arr = []string{
	ABS: "A",
	C:   "C",
}

func main() {
	fmt.Println("a: ", arr[ABS])
	fmt.Println("C: ", arr[C])
	fmt.Printf("arr %+v \n", arr)
	fmt.Printf("size: %+v \n", len(arr))
}