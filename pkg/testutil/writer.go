package testutil

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"testing"
)

type StringWriter struct {
	t               *testing.T
	receivedMessage string
}

func NewStringWriter(t *testing.T) *StringWriter {
	return &StringWriter{
		t: t,
	}
}

func (w *StringWriter) EXPECT(message string) {
	w.t.Cleanup(func() {
		if w.receivedMessage == "" {
			w.t.Errorf("expected message '%s' but got none", text.ANSI(message))
			return
		}
		if w.receivedMessage != message {
			w.t.Errorf("expected message '%s' but got '%s'", text.ANSI(message), text.ANSI(w.receivedMessage))
		}
	})
}

func (w *StringWriter) WriteString(s string) (n int, err error) {
	w.receivedMessage = s
	w.t.Log(text.ANSI(s))
	return len(w.receivedMessage), nil
}

func (w *StringWriter) Write(p *player.Player, s string) {
	w.receivedMessage = s
	w.t.Log(text.ANSI(s))
}
