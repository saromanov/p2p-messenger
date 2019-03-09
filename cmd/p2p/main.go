package main

import (
	"flag"

	"github.com/saromanov/p2p-messenger/internal/core"
	"github.com/saromanov/p2p-messenger/internal/server"
)

var (
	address = flag.String("address", "", "")
	name    = flag.String("name", "", "")
)

// start provides starting of the all stage on app
func start(address, name string) {
	app := core.New(name, address)
	app.Start()
	go func() {
		if err := server.New(address, app); err != nil {
			panic(err)
		}
	}()
	select {}
}
func main() {
	flag.Parse()
	if *address == "" {
		panic("address is not defined")
	}
	if *name == "" {
		panic("name is not defined")
	}

	start(*address, *name)
}
