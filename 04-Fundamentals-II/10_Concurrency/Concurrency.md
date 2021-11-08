# Concurrency

## Features

* Go has CSP (Communicating sequential processes) Style concurrency 
* See for details: https://en.wikipedia.org/wiki/Communicating_sequential_processes
* Communication between goroutines is handled via channels
* See main.go and benchmark.sh

### Simple Example

Each HTTP request is handled in a goroutine!

```golang
package main

import (
	"log"
	"net/http"
)

type server struct {}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", &server{}))
}

```


### Sleep
Channel example ([Run code](https://play.golang.org/p/F2jKWpWLJGP)):

Sleeping is possibility, but a poor man's approach

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


### Wait groups
Channel example ([Run code](https://play.golang.org/p/Whvt1Byyrwm)):

Waiting is much more suitable with a wait group

```golang
package main

import (
	"fmt"
	"sync"
)

func somethingOne(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("One")

}

func somethingTwo(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Two")
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go somethingOne(wg)
	go somethingOne(wg)

	wg.Wait()
}


```


### Context

Since golang has no threads, you will miss ThreadLocals...
Instead contexts are used and passed to each function.
Please read https://stackoverflow.com/questions/52799280/context-confusion-regarding-cancellation
for further details

See context/main.go



### Complex Example

([Run code](https://play.golang.org/p/tm_28R-r-MH)):
```golang
package main

import (
  "fmt"  // think about functionality and needed packages
  "time"
)

func readFromChannel(c chan int, done chan<- bool) {  // chan<- is send only
  for v := range c { // runs until channel is closed
    fmt.Println("Reading:", v)
  }
  done <- true // we are finished here
}

func Channel() {
  c := make(chan int,1)  // buffered channel, no locking for n elements
  done := make(chan bool) // non buffered channel, lock for each element
  // close the channel

  defer func() { <-done }() // read blocking done channel
  defer close(c)

  go readFromChannel(c, done) // start draining channel

  for i := 0; i < 200; i++ {
  	fmt.Println("Writing:", i)
	c <- i
  }

  time.Sleep(10*time.Second)

}
```
