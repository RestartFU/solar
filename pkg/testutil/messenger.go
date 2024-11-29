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

func (m *Messenger) EXPECT(messages ...string) {
	m.expectedMessages = append(m.expectedMessages, messages...)
	m.t.Cleanup(func() {
		if len(m.expectedMessages) != len(m.receivedMessages) {
			m.t.Errorf("expected %d messages, but got %d", len(m.expectedMessages), len(m.receivedMessages))
			return
		}

		for i, expected := range m.expectedMessages {
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
