# Variables

# Use it, or not define it

A variable in Go has to be used!

Unused Variables are compile errors!
This is an important feature of Go that makes Go Code more readable and maintainable.


## Inferred Variables

Variables are normally inferred.

```golang

var s = "test"

v := 42

```

## Define Variables with an explicit type

Go uses a reverse order for defining the name and the type.

It might look strange, after a few usages, it will turn into normality.


```golang

var name string
name = "test"

var value int
value = 33

```


# Pro Tip

a var definition can lead to poor performance because the desired value is 
somewhere in the memory and not in the processors register...

This litte example is slow, because the variable n is stored in memory and no
processor register is used for this simple function...

```golang
package main

func main() {
       f([]int{1,1})
}

var n =0

func f(s []int) int {
  n = 0
  for _,v := range s {
    n +=v
  }
  return n
}
```

You can use https://godbolt.org/ if you want to understand what assemby code the go compiler is producing
