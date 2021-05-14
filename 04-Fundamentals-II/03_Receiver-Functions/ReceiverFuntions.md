# Receiver Functions

In go, functions are not attached to anything. They can be called from everywhere unless they are public.
To attach functions to struct, the function needs a receiver.

```golang
type Point struct {
    X int
    Y int
}

func (p Point) GetX() int {
	retun p.X
}
```

The Receiver Function comes in to variations
- Value Receiver
- Pointer Receiver

## Value Receiver
A Value Receiver Function cannot change the internal state of the struct.

```golang
type Point struct {
    X int
    Y int
}

func (p Point) Coordinates() (int, int) {
    return p.X, p.Y
}

```

## Pointer Receiver
A Point Receiver can change the internal state of the struct.


```golang
type Point struct {
    X int
    Y int
}

func (p *Point) SetX(x int) {
	p.X = x
}
```

# Excercise

Correct the Test for setting the name of the struct with a receiver function.


# Pro Tips

## Location of receiver functions
A Receiver Function can be located in a different go-File, but has to be in the same package

This is used in the famous stringer generator to generate a Java toString() like behavior.

See https://pkg.go.dev/golang.org/x/tools/cmd/stringer for more information
