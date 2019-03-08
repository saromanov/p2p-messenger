// Package peer defines peer for the messenger
package peer

import (
	"net"

	"golang.org/x/crypto/ed25519"
)

// Peer defines peer for the messagener
type Peer struct {
	Name       string
	PublicKey  ed25519.PublicKey
	Connection net.Conn
}

// New creates a new peer
func New(c net.Conn) (*Peer, error) {
	return &Peer{
		Name:       "123",
		Connection: c,
	}, nil
}
