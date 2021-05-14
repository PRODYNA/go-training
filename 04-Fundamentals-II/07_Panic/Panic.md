# Panic

panic is an exception like element and stops the executing of the actual function and keeps
returning until a recover is handling the panic.

Normally the main function, started by the go runtime will handle the panic and stops your programm.

So, conclusion is, panic without recover is only for non-server programs acceptable.

```golang

func main() {
	
	...
	DoBusinessLogic()
	...
}

func DoBusinessLogic() {
	
	...
	panic("Not Expecting this")
}
```
