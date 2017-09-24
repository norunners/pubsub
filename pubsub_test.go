package pubsub

import "testing"

const (
	testTopic   = "testTopic"
	testMessage = "testMessage"
)

func TestPubSub_PubSub(t *testing.T) {
	t.Parallel()

	ps := New()
	rec := newMock()
	ps.Sub(rec, testTopic)

	expected := testMessage
	ps.Pub(expected, testTopic)
	actual := <-rec.ch
	if expected != actual {
		t.Errorf("Expected '%s' but found '%s'.", expected, actual)
	}
}

func TestPubSub_None(t *testing.T) {
	t.Parallel()

	ps := New()
	rec := newMock()
	ps.Sub(rec, testTopic)

	ps.Pub(testMessage, "none")
	if len(rec.ch) != 0 {
		t.Errorf("Nonsuscriber received a message.")
	}
}

func TestPubSub_UnSub(t *testing.T) {
	t.Parallel()

	ps := New()
	rec := newMock()
	us := ps.Sub(rec, testTopic)
	us()

	ps.Pub(testMessage, testTopic)
	if len(rec.ch) != 0 {
		t.Errorf("Unsubscriber received a message.")
	}
}

func TestPubSubPlus_UnSub(t *testing.T) {
	t.Parallel()

	ps := NewPlus()
	rec := newMock()
	ps.Sub(rec, testTopic)
	ps.UnSub(testTopic)

	ps.Pub(testMessage, testTopic)
	if len(rec.ch) != 0 {
		t.Errorf("Unsubscriber received a message.")
	}
}

func TestPubSubPlus_UnSubAll(t *testing.T) {
	t.Parallel()

	ps := NewPlus()
	rec := newMock()
	ps.Sub(rec, testTopic)
	ps.UnSubAll()

	ps.Pub(testMessage, testTopic)
	if len(rec.ch) != 0 {
		t.Errorf("Unsubscriber received a message.")
	}
}

type mockReceiver struct {
	ch chan string
}

func newMock() *mockReceiver {
	return &mockReceiver{ch: make(chan string)}
}

func (rec *mockReceiver) Receive(msg interface{}) {
	if val, ok := msg.(string); ok {
		rec.ch <- val
	}
}
