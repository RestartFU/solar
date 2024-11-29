package testutil

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"testing"
)

type Messenger struct {
	t                *testing.T
	expectedMessages []string
	receivedMessages []string
}

func NewMessenger(t *testing.T) *Messenger {
	return &Messenger{
		t: t,
	}
}

func (w *Messenger) EXPECT(messages ...string) {
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

func (w *Messenger) Message(_ *player.Player, s string) {
	w.receivedMessages = append(w.receivedMessages, s)
	w.t.Log(text.ANSI(s))
}
