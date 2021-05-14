# Interfaces

Interfaces are  collections of functions (receiver functions), assigned to a type.

Interfaces are never implemented like in an object-oriented language.
Go uses a concept named "Duck Typing", which is very flexible and can operate outside package scopes


## Simple Example 

```golang
type PersonAdapter interface {
  List(string) *Person
  Insert(p *Person)
}
```

## Ducktyping

Docktyping works even on builtin types like map.

```golang

type Header map[string]string

func (h *Header) Add(k, v string) {
	h[k] = v
}

...

h := NewHeader()
h["Accept"] = "application/json"

```

# Interface as parameters

Since Go has no generics, interfaces can be used instead.
In combination with type switch, a generic-like behavior can be implemented.

Please remember: value.(type) is a cast to the type "type"

```golang
func DoStuff(value interface) {

    switch t := value.(type) {
        case int64: {
          fmt.Println("Type is an integer:", t)
        }
    }
    
}
```


## Pro Tip

Interfaces and structs are the basis for structuring your services.

For Mocking and Testing purposes, it's very easy to replace functionality for unit tests. 
