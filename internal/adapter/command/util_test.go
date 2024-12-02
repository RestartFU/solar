package command_test

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/restartfu/solar/internal/adapter/command"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAllowerPlayer_Allow(t *testing.T) {
	allower := command.AllowerPlayer{}

	for _, tc := range []struct {
		name     string
		source   cmd.Source
		expected bool
	}{
		{
			name:     "player source is allowed",
			source:   &player.Player{},
			expected: true,
		},
		{
			name:     "non-player source is not allowed",
			source:   nil,
			expected: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, allower.Allow(tc.source))
		})
	}
}
