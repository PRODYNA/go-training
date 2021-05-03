package main

import (
	"fmt"
	"math"
)

type t string

func main() {

	x := 12

	math.Abs(float64(x))

	s := "test"

	for _, r := range s {
		fmt.Printf("%c", r)
	}

	st := "test"


	fmt.Printf("%s", t(st))

}
