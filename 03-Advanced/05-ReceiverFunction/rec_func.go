package _05_ReceiverFunction

type Person struct {
	Firstname string
	Lastname string
}


func (p Person) SetName(f,l string) {
	p.Firstname = f
	p.Lastname = l
}


