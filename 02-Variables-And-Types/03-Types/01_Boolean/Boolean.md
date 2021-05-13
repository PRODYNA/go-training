# Boolean

The boolean type in go works like (mostly) all other languages.
Please keep in mind that the bool type is a buitin type (a value) and cannot be nil


## bool type
```golang
b := true

if b {
	fmt.Println("So true")
}
```


## default value
```golang
b := true

if b {
	fmt.Println("So true")
}
```

# Pro Tips

To define a nillable boolean (or any other builtin type) you have to use a pointer

```golang
var b *bool
b = nil

...

b := new(bool)

```
