package _05_Structs

type Person struct {
	Name string
	Age  int
}

type Address struct {
	Street string
	Zip    string
}

type Car struct {
	Engine struct {
		Type string
	}
}
