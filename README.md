pubsub
======
[![GoDoc](https://godoc.org/github.com/norunners/pubsub?status.svg)](https://godoc.org/github.com/norunners/pubsub) [![Build Status](https://travis-ci.org/norunners/pubsub.svg?branch=master)](https://travis-ci.org/norunners/pubsub) [![codecov](https://codecov.io/gh/norunners/pubsub/branch/master/graph/badge.svg)](https://codecov.io/gh/norunners/pubsub) [![Go Report Card](https://goreportcard.com/badge/github.com/norunners/pubsub)](https://goreportcard.com/report/github.com/norunners/pubsub)

Package `pubsub` is a trivial publish subscribe library.

Install
-------
```bash
go get github.com/norunners/pubsub
```

PubSub
------
PubSub is a trivial publish subscribe interface.
##### New:
```go
// Create a new pubsub.
ps := pubsub.New()
```
##### Publish:
```go
// Publish a message to a topic.
ps.Pub("Hello World!", "greetings")
```
##### Subscribe:
```go
// Subscribe a receiver to a topic.
ps.Sub(rec, "greetings")
```
##### Unsubscribe:
```go
// Subscribe a receiver to a topic.
unSub := ps.Sub(rec, "greetings")
...
// Unsubscribe a receiver from a topic.
unSub()
```

Receiver
--------
Receiver is satisfied by subscribers to receive messages.
##### Receive:
```go
func (rec *rec) Receive(msg interface{}) {
    fmt.Printf("msg: %v\n", msg)
}
```
##### Type assertion:
```go
func (rec *rec) Receive(msg interface{}) {
    if val, ok := msg.(string); ok {
        fmt.Printf("val: %s\n", val)
    }
}
```
##### Channels:
```go
func (rec *rec) Receive(msg interface{}) {
        rec.ch <- val
}
```

PubSubPlus
----------
PubSubPlus is a trivial publish subscribe interface with unsubscribe methods.
##### New plus:
```go
// Create a new pubsub plus.
ps := pubsub.NewPlus()
```
##### Unsubscribe topic:
```go
// Unsubscribe all receivers from a topic.
ps.UnSub("greetings")
```
##### Unsubscribe all:
```go
// Unsubscribe all receivers from all topics.
ps.UnSubAll()
```

License
-------
* [MIT License](LICENSE)
