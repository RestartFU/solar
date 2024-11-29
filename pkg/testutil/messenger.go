package testutil

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"testing"
)

type Messenger struct {
	t                *testing.T
	receivedMessages []string
}

func NewMessenger(t *testing.T) *Messenger {
	return &Messenger{
		t: t,
	}
}

func (m *Messenger) EXPECT(expectedMessages ...string) {
	m.t.Cleanup(func() {
		if len(expectedMessages) != len(m.receivedMessages) {
			m.t.Errorf("expected %d messages, but got %d", len(expectedMessages), len(m.receivedMessages))
			return
		}

		for i, expected := range expectedMessages {
			received := m.receivedMessages[i]
			if received != expected {
				m.t.Errorf("expected message '%s' at index %d, but got '%s'", text.ANSI(expected), i, text.ANSI(received))
			}
		}
	})
}

func (m *Messenger) Message(_ *player.Player, s string) {
	m.receivedMessages = append(m.receivedMessages, s)
	m.t.Log(text.ANSI(s))
}
