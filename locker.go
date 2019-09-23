package peeridlocker

import (
	"sync"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/tevino/abool"
)

// PeerIDLocker is used to handle mutex lock/unlock
// on a per-peerID basis for systems that may need
// concurrent access to peerID specific resources
// such as IPNS record publishing
//
// The only time this blocks is when creating
// the initial lock, afterwards all writes for a record
// are non-blocking unless another record is first published
// at the same time.
type PeerIDLocker struct {
	locks map[peer.ID]*abool.AtomicBool
	mux   sync.RWMutex
}

// New returns a fresh instance of PeerIDLocker
func New() *PeerIDLocker {
	return &PeerIDLocker{
		locks: make(map[peer.ID]*abool.AtomicBool),
	}
}

// Create is like exists, except it
// populates the map if the entry does not exist
func (pl *PeerIDLocker) Create(pid peer.ID) {
	if !pl.Exists(pid) {
		pl.mux.Lock()
		pl.locks[pid] = abool.New()
		pl.mux.Unlock()
	}
}

// Exists check if we have a lock for this peerID
func (pl *PeerIDLocker) Exists(pid peer.ID) bool {
	pl.mux.RLock()
	_, exists := pl.locks[pid]
	pl.mux.RUnlock()
	return exists
}

// Lock obtains a lock for the peerID
func (pl *PeerIDLocker) Lock(pid peer.ID) {
	pl.Create(pid)
	pl.mux.RLock()
	pl.locks[pid].SetToIf(false, true)
	pl.mux.RUnlock()
}

// Unlock reverts the peerID lock
func (pl *PeerIDLocker) Unlock(pid peer.ID) {
	pl.Create(pid)
	pl.mux.RLock()
	pl.locks[pid].SetToIf(true, false)
	pl.mux.RUnlock()
}
