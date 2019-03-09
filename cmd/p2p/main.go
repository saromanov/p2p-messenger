package main

import (
	"flag"

	"github.com/saromanov/p2p-messenger/internal/core"
	"github.com/saromanov/p2p-messenger/internal/server"
)

var address = flag.String("address", "", "")

// start provides starting of the all stage on app
func start(address string) {
	go func() {
		if err := server.New(address); err != nil {
			panic(err)
		}
	}()

	app := core.New("message", address)
	app.Start()
	select {}
}
func main() {
	flag.Parse()
	if *address == "" {
		panic("address is not defined")
	}

	start(*address)
}
