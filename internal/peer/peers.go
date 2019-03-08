package peer

import "sync"

// Peers defines struct for store of map of peers
type Peers struct {
	mu    *sync.Mutex
	peers map[string]*Peer
}

// newPeers inits struct of peers
func newPeers() *Peers {
	return &Peers{
		mu:    &sync.Mutex{},
		peers: make(map[string]*Peer),
	}
}
