# Testing in Go

Go has a builtin testing library that covers most of the test functionally needed in daily programming.

There is a library [github.com/stretchr/testify](https://github.com/stretchr/testify) that extends the buitin testing functionaly with assert statements very similar to JUnit.

## Visibility in Go

Keep in mind that Go has only to types of visibility
* public ->  starts with an uppercase letter
* package/private -> starts with a lowercase letter

So everything within a package is visible to other Go files in the same package.


## Naming

A test in go is written as separate file, with the suffix filename_<test.go> for the gofile filename.go

```
fact.go
fact_test.go
```

The file process_test.go is a normal Go file and contains the tests.

Each Unit-Test is a function in this file and start with Test.

## Example
fact.go
```golang
package math

func Fact(i uint32) uint64 {
 ...
}

```
fact_test.go
```golang
package math

import (
 "testing"
)

func TestFact(t *testing.T) {
  if Fact(12) != 479001600 {
    t.Error("Failed for some reason")
    t.Fail()
  }
}
```

## Run the test

### Simple run
```go test```

### Run all tests

```go test - run ''```
### With coverage
```go test -coverprofile cp.out```

### View coverage results

In order to see the coverage `per function` of your project, run:

```
go tool cover -func coverage.out
```

the output will look like the following:

```
github.com/myorg/jetstarter-service/adapters/jfs_adapter.go:20:                 NewJFSAdapter                 100.0%
github.com/myorg/jetstarter-service/adapters/jfs_adapter.go:27:                 InitiateJFS                   0.0%
github.com/myorg/jetstarter-service/adapters/scheduling_service_adapter.go:22:	NewSchedulingServiceAdapter   100.0%
github.com/myorg/jetstarter-service/adapters/scheduling_service_adapter.go:29:	CreateSchedulingJob           82.4%
github.com/myorg/jetstarter-service/adapters/scheduling_service_adapter.go:56:	GetSchedulingJob			  82.4%
github.com/myorg/jetstarter-service/ports/handler.go:280:                       noAuthorizationHeader         0.0%
github.com/myorg/jetstarter-service/ports/handler.go:289:                       internalServerErrorResponse	  100.0%
                                                       total:                   (statements)                  62.4%
```

If you want to see the results in an `html` output in your browser, run the following command:

```
go tool cover -html coverage.out
```

A browser window will open and display the covered and uncovered lines of every file.

### Running only a subset of the tests
```go test -run Integration```

This would only call tests that are named IntegrationXXXTest

## Using testify

fact.go remains the same

fact_test.go
```golang
package math

import (
 "testing"
 "github.com/stretchr/testify/assert"
)

func TestFact(t *testing.T) {
  asssert.Equals(t, 479001600, Fact(12))
}
```

## Ordered/Nested Tests

Ordered and Nested Tests are very easy to create and they are not cursed like in the Java world.

```golang
func TestRunner(t *testing.T) {

  x := NewX(...)
	
  t.Run("First", func(t *testing.T) {
    // doing stuff on x
  })
  
  t.Run("Second", func(t *testing.T) {
    // doing more stuff on x
  })

}
```


# Pro Tips

Go has also a builtin functionality for benchmarks.
See https://golangdocs.com/benchmark-functions-in-golang for details

For Goland/IntelliJ: here is a cheat sheet for testing go https://www.jetbrains.com/help/go/performing-tests.html#run-tests-from-the-context-menu
