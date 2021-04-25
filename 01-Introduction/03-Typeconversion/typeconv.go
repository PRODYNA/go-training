package main

import (
	"fmt"
	"math"
)

func main() {

	x := 12

	math.Abs(float64(x))

	s := "test"

	for _, r := range s {
		fmt.Printf("%c", r)
	}
}
