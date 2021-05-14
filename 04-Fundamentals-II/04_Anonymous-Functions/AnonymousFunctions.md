# Anonymous functions

## Functions as variables

```golang
duplicate := func (s string) string {
  return s + s
}

fmt.Println(duplicate("Test"))

// Returns TestTest
```


## Returning a function


```golang
func main() {
 fmt.Println(duplicate("Test")())
}

func duplicate(s string) func() string {

  return func () string {
    return s + s
  }

}
```


## Funstions as parameters

```golang
func main() {

  duplicate := func(s string) string {
    return s + s
  }

  execute(duplicate, "Test")

}

func execute(f func(string) string, name string) {
  fmt.Println(f(name))
}

```


## Naming consequences

You can see that if the naming of parameters would be in order like java, the syntax would be unreadable

```golang
func execute(func(string) string f, string name) {
  fmt.Println(f(name))
}

```
