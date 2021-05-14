# Functions

## Scope

The scope of a function is determined by the first letter
* Uppercase \
  The function is public and everywhere available
* Lowercase \
  The function is private/package local and can be called only within the module/package

## Parameter

Functions can have multiple parameters, they are normally named with lowercase names.
If parameters are of the same type, they are grouped.

```golang

func Add(a,b int) int {
	return a+b
}
```

The last parameter can be an ellipse parameter.
Documentation has to be simple and on point. The first word must be the name of the function.

```golang
// Add sums all the values and stores them into a map with the given key.
func Add(key string, values ...int) int {
    ...
}
```

## Return types

Functions can have more than one return type.
This will be used later for the error handling.


```golang
// NewPerson creates a new Person.
func ListPersons(req PersonRequest ) *[]Person, error {
	...
    return &Person{Name = "Frank"}, nil
}
```
