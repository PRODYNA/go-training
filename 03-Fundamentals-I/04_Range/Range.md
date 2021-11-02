#  Range

The `range` form of the `for` loop iterates over a slice or map.

## Range over a slice

When ranging over a slice, two values are returned for each iteration. The first is the `index`, and the second is a copy of the element at that index.

For example ([Run code](https://play.golang.org/p/MoSqZXcLtkl)):
```go
package main

import "fmt"

var names = []string{"Frank", "Martin", "Chris"}

func main() {
	for i, v := range names {
		fmt.Printf("%d. %s writes Go!\n", (i + 1), v)
	}

	// Print the names again but omit the index
	for _, v := range names {
		fmt.Printf("%s writes Go!\n", v)
	}
}
```

## Range over a map

When ranging over a map, two values are returned for each iteration. The first is the `key` of the map and the second is the `value` of that `key`.

For example ([Run code](https://play.golang.org/p/TAiIwLs3hEJ)):
```golang
package main

import "fmt"

func main() {
	var machMap = map[string]float64{
		"F-16":   1.6,
		"SR-71":  3.2,
		"Mig-25": 2.83,
	}

	fmt.Println(machMap)

	for k, v := range machMap {
		fmt.Printf("%s can go up to mach: %.2f!\n", k, v)
	}

	// Print the keys again but omit the value
	fmt.Println("\nPrint only airplane names:")
	for k, _ := range machMap {
		fmt.Println(k)
	}
}
```
