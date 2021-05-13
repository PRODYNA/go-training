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

There is a function named "Init()" which is executed if you import a package.
E.g. the SQL Drivers are imported in this way.
