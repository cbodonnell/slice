package main

import (
	"fmt"

	"github.com/cheebz/slice"
)

func main() {
	a := slice.Insertable{1, 2, 4, 5}
	fmt.Println(a)
	value := 3
	index := 2
	a = a.InsertValueAtIndex(value, index)
	fmt.Println(a)
}
