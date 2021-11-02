
# Structs

In go, structs are used to structure data.

The fields and the struct itselfs are named like this

- **public** Uppercase first letter
- **private/package** Lowercase first letter

Example of a simple struct

```golang
type Point struct {
    X int
    Y int
}
```

Structs can contain any type other types, including pointers and of course structs.

```golang
type Triangle struct {
    A *Point
    B *Point
    C *Point
}

type Point struct {
    X int
    Y int
}
```

Structs can be nested. Be aware that you cannot reuse nested structs. 

```golang
type Car struct {
    Engine struct {
      Type string
    }
}
```

Creating an instance of a struct is simple done by using the name of the struct with curly braces.
Values can be passed by stating the name of the field and the value.

```golang
type Person struct {
	Name string
}

func CreatePerson(name string) Person {
	return Person{
	    Name: "PRODYNA Guy",	
    }
}
```


# Exercise

Enhance the struct 'Person' with the Address and fill the values, so the test runs properly.


# Pro Tips

## Object inheritance

It's possible to derive structs and achieve an object-oriented behavior.
By using the Identified struct inside Person without an identifier, all public Fields and Receiver-Functions
of the Identified struct are available in the Person struct.

```golang
type Identified struct {
    Id string
}

type Person struct {
    Identified
    Name string
}
```
