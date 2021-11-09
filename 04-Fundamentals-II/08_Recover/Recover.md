# Recover

A panic can be caught within a deferred function and with the function
recover(), the panic can be resolved or re-panic'd.


```golang
func main() {
	Recovering()
}


func Recovering() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	Panicking()
}

func Panicking() {
	panic(errors.New("One of these days..."))
}

```
