package peer

import (
	"errors"
	"sync"
)

var errPeerNotFound = errors.New("peeer is not found")

// Peers defines struct for store of map of peers
type Peers struct {
	mu    *sync.RWMutex
	peers map[string]*Peer
}

// newPeers inits struct of peers
func newPeers() *Peers {
	return &Peers{
		mu:    &sync.RWMutex{},
		peers: make(map[string]*Peer),
	}
}

// Add provides inserting of the new peer
func (p *Peers) Add(name string, peer *Peer) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.peers[name] = peer
}

// Get provides getting of the peer by the name
func (p *Peers) Get(name string) (*Peer, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	peer, ok := p.peers[name]
	if !ok {
		return nil, errPeerNotFound
	}
	return peer, nil
}
