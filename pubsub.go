// Package pubsub is a trivial publish subscribe library.
package pubsub

import "sync"

// PubSub is a trivial publish subscribe interface.
type PubSub interface {
	Pub(msg interface{}, topic string)
	Sub(rec Receiver, topic string) UnSub
}

// Receiver is satisfied by subscribers to receive messages.
// NOTE: Receive is called in a goroutine.
type Receiver interface {
	Receive(msg interface{})
}

// UnSub unsubscribes a receiver from a topic.
type UnSub func()

// pubSub satisfies PubSub.
type pubSub struct {
	mtx sync.RWMutex
	idx uint64
	reg map[string]map[uint64]Receiver
}

// New creates a new PubSub.
func New() PubSub {
	return newPubSub()
}

// newPubSub creates a new pubSub.
func newPubSub() *pubSub {
	reg := make(map[string]map[uint64]Receiver)
	return &pubSub{reg: reg}
}

// Pub publishes a message to all subscribers of a topic.
func (ps *pubSub) Pub(msg interface{}, topic string) {
	ps.mtx.RLock()
	defer ps.mtx.RUnlock()

	for _, rec := range ps.reg[topic] {
		go rec.Receive(msg)
	}
}

// Sub subscribes a receiver to a topic.
// Returns an unsubscribe callback function.
func (ps *pubSub) Sub(rec Receiver, topic string) UnSub {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if ps.reg[topic] == nil {
		ps.reg[topic] = make(map[uint64]Receiver)
	}

	idx := ps.idx
	ps.idx++
	ps.reg[topic][idx] = rec

	return func() {
		ps.unSub(idx, topic)
	}
}

// unSub unsubscribes a receiver from a topic.
func (ps *pubSub) unSub(idx uint64, topic string) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	delete(ps.reg[topic], idx)
}

// PubSubPlus is a trivial publish subscribe interface with unsubscribe methods.
type PubSubPlus interface {
	PubSub
	UnSub(topic string)
	UnSubAll()
}

// pubSubPlus satisfies PubSubPlus.
type pubSubPlus struct {
	*pubSub
}

// NewPlus creates a new PubSubPlus.
func NewPlus() PubSubPlus {
	ps := newPubSub()
	return &pubSubPlus{pubSub: ps}
}

// UnSub unsubscribes all receivers from a topic.
func (ps *pubSubPlus) UnSub(topic string) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	delete(ps.reg, topic)
}

// UnSubAll unsubscribes all receivers from all topics.
func (ps *pubSubPlus) UnSubAll() {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	for topic := range ps.reg {
		delete(ps.reg, topic)
	}
}
