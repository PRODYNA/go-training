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
