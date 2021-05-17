# Conditionals

Go offers flow control with the conditional statements `if` and `switch`.

## If

Go's `if` statements are like its `for` loops; the expression need not be surrounded by parentheses `( )` but the braces `{ }` are required. 

For example ([Run code](https://play.golang.org/p/oxiKpLPSidg)):

```go
package main

import (
	"fmt"
)

func main() {
	bob := "PRODYNA"

	if bob == "PRODYNA" {
		fmt.Println("Hi " + bob + " colleague!")
	} else {
		fmt.Println("Hi stranger!")
	}
}
```

## If with a short statement

The `if` statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the if.

For example ([Run code](https://play.golang.org/p/Za8qekglXIj)):

```golang
package main

import (
	"fmt"
	"math"
)

func isPositive(n float64) bool {
	return n > 0
}

func main() {
	num := -10.2

	if v := math.Abs(num); isPositive(v) == true {
		fmt.Println("The number is positive")
	} else {
		fmt.Println("The number is negative")
	}
}
```
