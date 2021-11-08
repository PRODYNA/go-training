## Channels


Channels are message-like types for communication between goroutines.

* Channels are typed like arrays
* They can be unidirectional or bidirectional and readable/writeable
* There is an internal capacity wich acts like a buffer
* Channels can (have to) be closed
* Pointers on channels are not needed
* Channels need to be handled correctly, panics and deadlocks are NORMAL while development


### Simple Example
Channel example ([Run code](https://play.golang.org/p/g6TmELvmdhQ)):
```golang
package main

import (
"fmt"
)

func main() {

	c := make(chan string,1)
	c <- "test"
	fmt.Println(<-c)
	close(c)
}

```

### Sending only channels

Channel example ([Run code](https://play.golang.org/p/6S9uT8Tl3jZ)):
```golang
package main

import (
    "fmt"
)

func main() {

	c := make(chan string,1)
	send(c)
	fmt.Println(<-c)
	close(c)
}


func send(c chan<- string) {
   c <- "test"
}
```


### Write Example
Channel example ([Run code](https://play.golang.org/p/bAn2vcvfyGY)):
```golang
package main

import (
	"fmt"
)

func main() {

	c := make(chan string,1)
	defer close(c)
	
	c <- "Test"
	read(c)

}


func read(c <-chan string) {
   fmt.Println(<-c)
}

```


### Range Example
Channel example ([Run code](https://play.golang.org/p/F2jKWpWLJGP)):
```golang
package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string,1)
	defer close(c)
	
	go read(c)
	
	c <- "Test"
	c <- "Test2"
	
	time.Sleep(10 * time.Second)
}


func read(c <-chan string) {
   for v := range c {
     fmt.Println(v)
   }
}
```


### Switch Example
Channel example ([Run code](https://play.golang.org/p/nozeGa8eR1P)):
```golang
package main

import (
	"fmt"
)

func main() {

	c := make(chan interface{},10)
	defer close(c)
	
	
	c <- true
	c <- "Test"
	
	read(c)
	read(c)

}


func read(c <-chan interface{}) {
    i := <-c
    switch i.(type) {
    case bool:
	fmt.Println("Got a boolean")
    case string:
	fmt.Println("A String, my friend")
    default:
	fmt.Println("Something different")
    }
}
```
