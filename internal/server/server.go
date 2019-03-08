// Package server provides definition on the TCP server
package server

import (
	"log"
	"net"
	"syscall"

	"github.com/pkg/errors"
)

// New provides making of TCP server
func New(address string) error {
	l, err := net.Listen("tcp", address)
	if err != nil {
		return errors.Wrap(err, "unable to listen TCP")
	}

	log.Printf("Server started on %s\n\n", address)
	for {
		conn, err := l.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				log.Printf("accept temp err: %v", ne)
				continue
			}

			log.Printf("accept error: %v", err)
			return err
		}
		go handle(conn)
	}
	return nil
}

func handle(conn net.Conn) {

}

func setLimit() {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}

	log.Printf("set cur limit: %d", rLimit.Cur)
}
