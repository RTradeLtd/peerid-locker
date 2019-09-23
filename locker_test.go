package peeridlocker

import (
	"sync"
	"testing"
	"time"

	ci "github.com/libp2p/go-libp2p-core/crypto"
	peer "github.com/libp2p/go-libp2p-core/peer"
)

// basic consistency test
func TestPidLocker(t *testing.T) {
	_, pubKey, err := ci.GenerateKeyPair(ci.Ed25519, 256)
	if err != nil {
		t.Fatal(err)
	}
	pid, err := peer.IDFromPublicKey(pubKey)
	if err != nil {
		t.Fatal(err)
	}
	locker := New()

	// test exists
	if locker.Exists(pid) {
		t.Fatal("pid lock should not exist")
	}

	// test create
	locker.Create(pid)

	// test exists
	if !locker.Exists(pid) {
		t.Fatal("pid lock should exist")
	}

	// test lock
	locker.Lock(pid)

	// test unlock
	locker.Unlock(pid)
}

// attempts to trigger race conditions
// with repeated execution with the same peerID
func TestPidLock_Race_Single_Peer(t *testing.T) {
	locker := New()
	wg := &sync.WaitGroup{}
	testFunc := func(pid peer.ID, waitTime time.Duration) {
		defer wg.Done()
		time.Sleep(waitTime)
		// test create
		locker.Create(pid)

		// test lock
		locker.Lock(pid)

		// test exists
		locker.Exists(pid)

		// test unlock
		locker.Unlock(pid)
	}
	_, pubKey, err := ci.GenerateKeyPair(ci.Ed25519, 256)
	if err != nil {
		t.Fatal(err)
	}
	pid, err := peer.IDFromPublicKey(pubKey)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go testFunc(pid, time.Nanosecond+time.Duration(i))
		wg.Add(1)
		go testFunc(pid, 0)
	}
	wg.Wait()
}

// attempt to trigger race conditions
// with execution of randomly generated peerIDs.
// the long sleep times are to ensure we queue
// up requests despite the wait time required
//for randomly generated the keys
func TestPidLock_Race_Many_Peer(t *testing.T) {
	locker := New()
	wg := &sync.WaitGroup{}
	testFunc := func(pid peer.ID, waitTime time.Duration) {
		defer wg.Done()
		time.Sleep(waitTime)
		// test create
		locker.Create(pid)

		// test lock
		locker.Lock(pid)

		// test exists
		locker.Exists(pid)

		// test unlock
		locker.Unlock(pid)
	}
	for i := 0; i < 500; i++ {
		_, pubKey, err := ci.GenerateKeyPair(ci.Ed25519, 256)
		if err != nil {
			t.Fatal(err)
		}
		pid, err := peer.IDFromPublicKey(pubKey)
		if err != nil {
			t.Fatal(err)
		}
		wg.Add(1)
		go testFunc(pid, time.Nanosecond*10)
		wg.Add(1)
		go testFunc(pid, time.Nanosecond*9)
		wg.Add(1)
		go testFunc(pid, time.Nanosecond*8)
		wg.Add(1)
		go testFunc(pid, time.Nanosecond*7)
	}
	wg.Wait()
}
