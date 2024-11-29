package testutil

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"testing"
)

type StringWriter struct {
	t                *testing.T
	expectedMessages []string
	receivedMessages []string
}

func NewStringWriter(t *testing.T) *StringWriter {
	return &StringWriter{
		t: t,
	}
}

// EXPECT allows expecting multiple messages in order.
func (w *StringWriter) EXPECT(messages ...string) {
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

// WriteString records the received message.
func (w *StringWriter) WriteString(s string) (n int, err error) {
	w.receivedMessages = append(w.receivedMessages, s)
	w.t.Log(text.ANSI(s))
	return len(s), nil
}

// Write records the received message for a specific player.
func (w *StringWriter) Write(_ *player.Player, s string) {
	w.receivedMessages = append(w.receivedMessages, s)
	w.t.Log(text.ANSI(s))
}
