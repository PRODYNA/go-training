# Loops

## For

Go has only one looping construct, the `for` loop.

The basic `for` loop has three components separated by semicolons:

- the init statement: executed before the first iteration
- the condition expression: evaluated before every iteration
- the post statement: executed at the end of every iteration

The init statement will often be a short variable declaration, and the
variables declared there are visible only in the scope of the `for`
statement.

The loop will stop iterating once the boolean condition evaluates to `false`.

For example ([Run code](https://play.golang.org/p/Yh8jRtIdbuT)):
```golang
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

**Note:** Unlike other languages like C, Java, or JavaScript there are no parentheses
surrounding the three components of the `for` statement and the braces `{}` are
always required.

The init and post statements are *optional*.

For example ([Run code](https://play.golang.org/p/iLbpPmG0di5)):

```golang
package main

import "fmt"

func main() {
	sum := 1
    // Check the fact that we do not initialise
    // the sum inside for loop.
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
```

## Where is "while" or "Do-while"?

There is no `while` or `do - while`. Go has only `for` which can act like `while` or `do - while`.

For example ([Run code](https://play.golang.org/p/-8ihiw0pxlX)):

```golang
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```

## Infinite loop

If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.

```golang
package main

func main() {
	for {
	}
}
```
