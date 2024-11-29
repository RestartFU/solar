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

func (s *Subscriber) EXPECT(messages ...string) {
	s.expectedMessages = append(s.expectedMessages, messages...)
	s.t.Cleanup(func() {
		if len(s.expectedMessages) != len(s.receivedMessages) {
			s.t.Errorf("expected %d messages, but got %d", len(s.expectedMessages), len(s.receivedMessages))
			return
		}

		for i, expected := range s.expectedMessages {
			received := s.receivedMessages[i]
			if received != expected {
				s.t.Errorf("expected message '%s' at index %d, but got '%s'", text.ANSI(expected), i, text.ANSI(received))
			}
		}
	})
}

func (s *Subscriber) UUID() uuid.UUID {
	return s.id
}

func (s *Subscriber) Message(a ...any) {
	str := fmt.Sprint(a...)
	s.receivedMessages = append(s.receivedMessages, str)
	s.t.Log(text.ANSI(s))
}
