package main

import "flag"

var address = flag.String("address", "", "")

func main() {
	flag.Parse()
	if *address == "" {
		panic("address is not defined")
	}
}
