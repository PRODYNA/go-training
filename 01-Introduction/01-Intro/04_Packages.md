# Packages


Go programs using packages to divide Go Code into maintainable pieces.

There are some difference between Go and Java Packages

* Packages should be named with modules as existing URLs
* Only the last path element of a package is the qualifier
* Packages can be imported with an underscore without using anything (Init function is important) of it
* Unused Imports are compile-errors
* Packages can be renamed if there are duplicate names
  * In the import statement in the go file
  * In the go.mod file  
* Packages don't need a comma or semicolon for the import  


```golang

import (
	f "fmt"
	_ "strconv"
	"github.com/prj/aa/bb/cc"
)

func do() {
	
	cc.DoSomething()
	
	f.Println("Test")
}

```

## Pro Tip

There is a function named "init()" which is executed if you import a package.
E.g. the SQL Drivers are imported in this way.


## Run in Playground
For example ([Run code](https://play.golang.org/p/LO1l8Novpju)):


## go mod is powerful

You can use a package that does not exists in a repository
```golang
package main

import (
  "prodyna.com/go-lib/print"
)

func main() {
  print.PrintHello()
}

```

by simply replacing the required package by the real package

```
module dev.azure.com/frankratschinski/frankratschinski/_git/go-client
go 1.16
replace (
  prodyna.com/go-lib v1.0.1 => dev.azure.com/frankratschinski/frankratschinski/_git/go-lib v1.0.1-0.20211013100434-eeb0c549d957
)

require (
  prodyna.com/go-lib v1.0.1
)
```
