package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/restartfu/solar/internal/core/message"
	"github.com/restartfu/solar/internal/ports"
)

var Messenger ports.Messenger = message.Messenger{}
var Subscriber chat.Subscriber = chat.StdoutSubscriber{}

type playerAllower struct{}

func (playerAllower) Allow(src cmd.Source) bool {
	_, ok := src.(*player.Player)
	return ok
}
