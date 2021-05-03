# Basic Tools

## Go Mod

* Go mod is a tool to create a Module for Go.
* Each module has a "go.mod" file
* The "go.mod" file contain 
  * Module name
  * Go version
  * Dependencies
   

## Example Go Mod File
```
module github.com/prodyna/exciting-dashboard-service

go 1.13

require (
	github.com/denisenkom/go-mssqldb v0.10.0
	github.com/influxdata/influxdb-client-go/v2 v2.2.3
	github.com/prodyna/go-microservice-base v1.0.4
)

```
  

Usage ```go mod init "github.com/prodyna/go-training/chapter1```

See https://golang.org/doc/modules/gomod-ref for more details

## Go Build

Go build is the go compiler tool. It creates an executable binary.
There is no need for a build system like maven, make or other tools.
Later it is useful to put everything into a Docker File.

### Environment Variables

* GOOS - Go Operating System
* GOARCH - Go Architecture - The CPU
* GOROOT - Go Root DEPRECATED! Similar to classpath of Java

Usage ```go build cmd/main.go -o mygoprogram```

## Go Run

You can use the Go compiler as "interpreter" - it compiles and runs the go program on the fly.

Usage ```go run cmd/main.go```

## Go Test

Testing a Go program is simple.

Usage ```go test```


See https://golang.org/pkg/testing/ for more details

## IDE Support

You can use Idea, Goland (Special Version of Idea) or Vs-Code.

Using Vi, Emacs etc. is also possible :-)

Compiling Go on the console is faster than using an IDE.


