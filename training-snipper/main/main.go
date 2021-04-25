package main

import (
	"flag"
	"github.com/PRODYNA/go-training/snipper/snip"
)

func main() {
	out := flag.String("o", "", "where should the snipped sources go?")
	in := flag.String("i", "", "where should the snipped sources go?")

	flag.Parse()
	valid := true
	if *out == "" || *in == "" {
		valid = false
	}

	if !valid {
		flag.Usage()
	}

	s := snip.New(*out, *in)
	s.Snip()
}
