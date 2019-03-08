// Package server provides definition on the TCP server
package server

import (
	"log"
	"net"
	"syscall"

	"github.com/pkg/errors"
	"github.com/saromanov/p2p-messenger/internal/peer"
)

// New provides making of TCP server
func New(address string) error {
	setLimit()
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
	defer func() {
		conn.Close()
	}()

	log.Printf("New connection: %s", conn.RemoteAddr().String())

	_, err := peer.New(conn)
	if err != nil {
		panic(err)
	}
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
