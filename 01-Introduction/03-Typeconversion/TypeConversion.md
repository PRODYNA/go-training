# Type Conversion

Type conversion in go is done by built in functions which are named after the desired type


```golang

x := 12

math.Abs(float64(x))

```


## Converting Integers and Strings

For converting that can't be done with a cast, the strcnv package has useful functions.

```golang

v := strconv.Itoa(10)

// or

i,_ := strconv.Atoi(v)

```



# Pro Tips

Defining a new type also defines a conversion function 

```golang

type t string

st := "test"


fmt.Printf("%s", t(st))


```
