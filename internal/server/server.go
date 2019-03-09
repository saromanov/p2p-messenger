// Package server provides definition on the TCP server
package server

import (
	"fmt"
	"net"
	"syscall"

	"github.com/pkg/errors"
	"github.com/saromanov/p2p-messenger/internal/log"
	"github.com/saromanov/p2p-messenger/internal/peer"
	"github.com/saromanov/p2p-messenger/internal/core"
)

// New provides making of TCP server
func New(address string, cor *core.Core) error {
	setLimit()
	l, err := net.Listen("tcp", address)
	if err != nil {
		return errors.Wrap(err, "unable to listen TCP")
	}

	log.Infof("Server started on %s\n\n", address)
	for {
		conn, err := l.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				log.Infof("accept temp err: %v", ne)
				continue
			}

			log.Infof("accept error: %v", err)
			return err
		}
		go handle(conn, cor)
	}
	return nil
}

func handle(conn net.Conn, cor *core.Core) {
	defer func() {
		conn.Close()
	}()

	log.Infof("New connection: %s", conn.RemoteAddr().String())
	fmt.Println(conn)
	_, err := peer.New(conn)
	if err != nil {
		panic(err)
	}
	cor.Handle(conn)
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

	log.Infof("set cur limit: %d", rLimit.Cur)
}
