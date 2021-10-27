# Arrays

In Go, an array is a numbered sequence of elements of a specific length.

The type `[n]T` is an array of `n` values of type `T`.

The expression

```go
var tmpArray [10]int
```

declares a variable `tmpArray` as an array of ten integers.

An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays. 

For example ([Run code](https://play.golang.org/p/yp4ZO8qXrjw)):
```go
package main

import "fmt"

func main() {
	var a [4]string
	a[0] = "Hello"
	a[1] = "PRODYNA"
	a[2] = "Golang"
	a[3] = "developer"
	
	fmt.Println(a[0], a[1])
	fmt.Println(a)

}
```

# Slices

Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.

Unlike arrays, slices are typed only by the elements they contain (not the number of elements).

For example ([Run code](https://play.golang.org/p/3_04CaXl299)):
```go
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
```

## Slice defaults

When slicing, you may omit the high or low bounds to use their defaults instead. The default is zero for the low bound and the length of the slice for the high bound.

For the array

```go
var a [10]int
```
these slice expressions are equivalent:

```go
a[0:10]
a[:10]
a[0:]
a[:]
```

## Slice length and capacity

A slice has both a *length* and a *capacity*.

* The *length* of a slice is the number of elements it contains.
* The *capacity* of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The length and capacity of a slice s can be obtained using the expressions `len(s)` and `cap(s)`.

You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.

For example ([Run code](https://play.golang.org/p/kUhSiGF6SId)):
```go
package main

import "fmt"

func main() {
	s := []string{"verba", "volant", "scripta", "manent"}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

## Slices of slices

Slices can contain any type, including other slices. It can be seen as an equivalent of multidimentional arrays in Java.

For example ([Run code](https://play.golang.org/p/hozbc5GcK1Y)):
```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create an Arrow
	arrow:= [][]string{
		[]string{" ", " ", " ", " ", " ", " ", "*"},
		[]string{" ", " ", " ", " ", " ", "*", "*", "*"},
		[]string{" ", " ", " ", " ", "*", "*", "*", "*", "*"},
		[]string{" ", " ", " ", "*", "*", "*", "*", "*", "*", "*"},
		[]string{" ", " ", "*", "*", "*", "*", "*", "*", "*", "*", "*"},
		[]string{" ", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*"},
		[]string{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*"},
		[]string{" ", " ", " ", " ", " ", "|", "|", "|"},
		[]string{" ", " ", " ", " ", " ", "|", "|", "|"},
		[]string{" ", " ", " ", " ", " ", "|", "|", "|"},
	}

	for i := 0; i < len(arrow); i++ {
		fmt.Printf("%s\n", strings.Join(arrow[i], " "))
	}
}
```