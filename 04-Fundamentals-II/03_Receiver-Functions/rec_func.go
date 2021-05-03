package _3_Receiver_Functions

type Person struct {
	Firstname string
	Lastname string
}


func (p Person) SetName(f,l string) {
	p.Firstname = f
	p.Lastname = l
}


