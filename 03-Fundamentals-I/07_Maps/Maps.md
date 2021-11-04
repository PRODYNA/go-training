# Maps

A map maps keys to values.

The zero value of a map is `nil`.
A `nil` map has no keys, nor can keys be added.

The `make` function returns a map of the given type,
initialized and ready for use.

For example ([Run code](https://play.golang.org/p/7krybjXvPEn)):
```golang
package main

import "fmt"

type Person struct {
	Name        string
	Surname     string
	YearOfBirth int
}

var m map[string]Person

func main() {
	m = make(map[string]Person)
	m["ltorvalds"] = Person{
		"Linus", "Torvalds", 1969,
	}
	m["dritchie"] = Person{
		"Dennis", "Ritchie", 1941,
	}
	m["rpike"] = Person{
		"Rob", "Pike", 1956,
	}

	fmt.Println(m["ltorvalds"])
	fmt.Println(m["dritchie"])
}
```

## Map literals

Map literals are like struct literals, but the keys are required.


For example ([Run code](https://play.golang.org/p/r12-VvXihx6)):
```golang
package main

import "fmt"

type Person struct {
	Name        string
	Surname     string
	YearOfBirth int
}

var m map[string]Person

func main() {
	var mapByUsername = map[string]Person{
		"ltorvalds": Person{
			"Linus", "Torvalds", 1969,
		},
		"dritchie": Person{
			"Dennis", "Ritchie", 1941,
		},
		"rpike": Person{
			"Rob", "Pike", 1956,
		},
	}

	fmt.Println(mapByUsername ["ltorvalds"])
	fmt.Println(mapByUsername ["rpike"])
}
```

If the top-level type is just a type name, you can omit it from the elements of the literal.


For example ([Run code](https://play.golang.org/p/lO_6jXkpL_8)):
```golang
package main

import "fmt"

type Person struct {
	Name        string
	Surname     string
	YearOfBirth int
}

var m map[string]Person

func main() {
	var mapByUsername = map[string]Person{
		"ltorvalds": {
			"Linus", "Torvalds", 1969,
		},
		"dritchie": {
			"Dennis", "Ritchie", 1941,
		},
		"rpike": {
			"Rob", "Pike", 1956,
		},
	}

	fmt.Println(mapByUsername)
}
```

## Mutating Maps

Insert or update an element in map `m`:

```golang
m[key] = elem
```

Retrieve an element:

```golang
elem = m[key]
```

Delete an element:

```golang
delete(m, key)
```

Test that a key is present with a two-value assignment:

```golang
elem, ok = m[key]
```

```golang
if val, ok := dict["foo"]; ok {
    //do something here
}
```

If `key` is in `m`, `ok` is `true`. If not, `ok` is `false`.

If `key` is not in the map, then `elem` is the zero value for the map's element type.

*Note:* If `elem` or `ok` have not yet been declared you could use a short declaration form:

```golang
elem, ok := m[key]
```

Example ([Run code](https://play.golang.org/p/k1UHhiBIesO)):

```golang
package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name        string
	Surname     string
	YearOfBirth int
}

var m map[string]Person

func main() {
	var mapByUsername = map[string]Person{
		"ltorvalds": {
			"Linus", "Torvalds", 1969,
		},
		"dritchie": {
			"Dennis", "Ritchie", 1941,
		},
		"rpike": {
			"Rob", "Pike", 1956,
		},
	}

	// Test that a key is present with a two-value assignment
	_, aturingExists := mapByUsername["aturing"]

	fmt.Println("Does aturing exist? ", strconv.FormatBool(aturingExists))

	// Delete Linus Torvalds from map
	delete(mapByUsername, "ltorvalds")

	// Add Alan Turing to the map
	mapByUsername["aturing"] = Person{
		Name:        "Alan",
		Surname:     "Turing",
		YearOfBirth: 1912,
	}

	// Print map again
	fmt.Println(mapByUsername)
}
```
