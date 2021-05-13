# Pointers

Go uses pointers to distinguish between values and references to values.

## Behavior

* A pointer is a non-changeable reference to a value.
* You can't modify the value of the pointer with pointer-arithmetic

## Create a pointer from a value
```golang
s := "test"

p := &s
```


## De-Reference the pointer
You can get a pointer for any value

```golang
s := "test"

p := &s

fmt.Println(*s)
```

## Pass a pointer to a function
Pointers can be passed to functions
```golang
s := "test"
PrintString(&s)
...

func PrintString(s *string) {
   fmt.Println(*s)
}
```


## Pointing to nothing

Go has a value for an unassigned pointer: nil

```golang
var x *string

fmt.Println(x)
```


# Memory Model

## Values
* Values are normally stored in the stack and therefore very fast.
* Values can be still modified. But it results in a copy.
* A value is not garbage collected because it will be cleared when the stack frame is destroyed 

## Pointers
* A Pointer references a value that is stored in the heap
* When pointers are not longer used, they need to be garbage collected



# Pro Tips

Interfaces are always a pointer, but not marked visibly in the go code!

Functions that starts with New should always return a pointer or an interface.
Pointers to Interfaces are pointer to pointers! Don't use it.

## The new function
The function new creates a pointer to a buitin type
```golang
s := new(string)
```

## References
* https://golang.org/ref/mem
