package _8_Structs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStruct(t *testing.T) {

	p := Person{
		Name: "Frank",
		Age:  48,
	}

	assert.Equal(t, "{Frank 48 {Street 1234}}", fmt.Sprintf("%v", p))
}
