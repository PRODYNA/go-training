# Builtin functions

Go has some builtin functions to make programming easy when no generics are available

## make

The make function creates instances of complex types and return a pointer of the instance

```golang
m := make(map[string]*Request)

l := make(list[string],0)
```


## append

The append function adds an element to the end of a slice.
Keep in mind that you create a copy!

```golang
l := make([]string,0)
l = append(l, "test")
fmt.Println(l)
```


## len

len is used to get the amount/length of a slice
```golang
l := make([]string,0)
fmt.Println(len(l))
```

## cap

cap is used to get the amount of possible elements in a slice

```golang
l := make([]string,0,100)
fmt.Println(cap(l))
```
So you can append 100 values to the slice before the slice needs to be copied.
Thats very useful when you deal with huge slices.

## new

The new function creates a new instance of a value and return a pointer to the value

```golang
s := new(string)
```

Keep in mind that the value is not initialized with some values


## delete

Delete removes a value from a map
```golang
m := make(map[string]string)
m["x"] = "y"
delete(m, "x")
```


## copy

Copy copies values from a slice to another slice.
Keep in mind that the definition of source and destination is 
```golang
func copy(dst, src []Type) int
```


```golang
a := []string{"x"}
b := []string{"y"}
n := copy(b, a)
// b is now b,a
// n is now 1
```

## close

The close function acts only on channels, marking them as closed.
There is no autoclose in Go, the concept of defer is more powerful.

```golang
c := make(chan(string),1)
close(c)
```

