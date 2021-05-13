# Strings and Runes

Go has a built in type for strings. A string in Go cannot be null.
A string consists of runes, a funny name for characters...


```golang
s := "test"
fmt.Println(s)
```


## String Comparison

String comparison in go is easy - as long you do not have a pointer to a string


```golang

s := "test"

if s == "test" {
	fmt.Println("OK")
}


```


## Appending

Appending strings in go is done by the + operator.
You can use buffers for better performance

```golang

var name string

name = "Go"
name = name + " is cool"

```

## Iterating over characters (runes) 

Converting the rune with Printf
```golang

s := "test"

for _, r := range s {
    fmt.Printf("%c", r)
}

```

Converting the rune with the string() type conversion
```golang

s := "test"

for _, r := range s {
    fmt.Print(string(r))
}

```

## Pro Tips

* A rune in go is nothing special, it's a value of type int32

See bultin.go

```golang

type rune = int32

```


