package testutil

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"testing"
)

type Subscriber struct {
	t                *testing.T
	id               uuid.UUID
	expectedMessages []string
	receivedMessages []string
}

func NewSubscriber(t *testing.T) *Subscriber {
	return &Subscriber{
		t:  t,
		id: uuid.New(),
	}
}

func (w *Subscriber) EXPECT(messages ...string) {
	w.expectedMessages = append(w.expectedMessages, messages...)
	w.t.Cleanup(func() {
		if len(w.expectedMessages) != len(w.receivedMessages) {
			w.t.Errorf("expected %d messages, but got %d", len(w.expectedMessages), len(w.receivedMessages))
			return
		}

		for i, expected := range w.expectedMessages {
			received := w.receivedMessages[i]
			if received != expected {
				w.t.Errorf("expected message '%s' at index %d, but got '%s'", text.ANSI(expected), i, text.ANSI(received))
			}
		}
	})
}

func (w *Subscriber) UUID() uuid.UUID {
	return w.id
}

func (w *Subscriber) Message(a ...any) {
	s := fmt.Sprint(a...)
	w.receivedMessages = append(w.receivedMessages, s)
	w.t.Log(text.ANSI(s))
}
