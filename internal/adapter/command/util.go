package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/restartfu/solar/internal/core/message"
	"github.com/restartfu/solar/internal/ports"
)

var Messenger ports.Messenger = message.Messenger{}

type playerAllower struct{}

func (playerAllower) Allow(src cmd.Source) bool {
	_, ok := src.(*player.Player)
	return ok
}
