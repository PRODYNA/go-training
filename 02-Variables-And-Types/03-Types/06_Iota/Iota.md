# Iota

Iota is a special keyword for defining a sequence of numbers, starting with Zero.


```golang
const (
    Red   = iota
    Blue
    Green
)
```


## Enumerations are missing

Go does not have enumerations!
But the type system of Go is very powerful.


```golang

type Color int

const (
    Red Color = iota
    Blue
    Green
)
```

## But Go has no methods...


```golang

type Color int

const (
    Red Color = iota
    Blue
    Green
)

func (c Color) String() string {
    return [...]string{"Red", "Blue", "Green"}[c] 
}
```
