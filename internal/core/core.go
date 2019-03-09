// Package core defines main logic for the messenger
package core

import (
	"bufio"
	"net"
	"strings"

	"github.com/saromanov/p2p-messenger/internal/peer"
	"golang.org/x/crypto/ed25519"
)

// Core defines main logic
type Core struct {
	peers  *peer.Peers
	Name   string
	Addr   string
	PubKey ed25519.PublicKey
}

// New creates core logic
func New(name, addr string) *Core {
	return &Core{
		Name: name,
		Addr: addr,
	}
}

// Start provides initialization of the core
func (c *Core) Start() error {
	return nil
}

// Handle provides handling of incoming messages
func (c *Core) Handle(conn net.Conn) error {
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	scanr := bufio.NewScanner(r)
	for {
		scanned := scanr.Scan()
		if !scanned {
			if err := scanr.Err(); err != nil {
				return err
			}
			break
		}
		w.WriteString(strings.ToUpper(scanr.Text()) + "\n")
		w.Flush()
	}
	return nil
}
