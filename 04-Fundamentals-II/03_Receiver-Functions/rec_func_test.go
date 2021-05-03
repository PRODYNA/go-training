package _3_Receiver_Functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReceiver(t *testing.T) {

	p := Person{
		Firstname: "PRODYNA",
		Lastname:  "GUY",
	}

	p.SetName("Golang" ,"Guy")
	assert.Equal(t, "Golang", p.Firstname)
}
