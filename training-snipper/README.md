This little tool can be used to copy the complete directory structure whilemoving snippets that the training attendees should enter themselves.
E.g.

```go
package main

import "log"

func main() {
	// SNIP START
	log.Printf("This line will be removed.")
	// SNIP END
	log.Printf("This line will not be removed.")
}
``` 
This little tool can be used to copy the complete directory structure whilemoving snippets that the training attendees should enter themselves.
will be transformed to 

```go
package main

import "log"

func main() {
	log.Printf("This line will not be removed.")
}
``` 
