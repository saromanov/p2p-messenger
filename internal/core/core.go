// Package core defines main logic for the messenger
package core

import (
	"net"

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
	return nil
}
