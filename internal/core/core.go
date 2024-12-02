package core

import (
	"github.com/restartfu/solar/internal/core/domain/message"
	"github.com/restartfu/solar/internal/core/ports"
	"strings"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/world"
)

var (
	Messenger      ports.Messenger      = message.Messenger{}
	Subscriber     chat.Subscriber      = chat.StdoutSubscriber{}
	UserRepository ports.UserRepository = nil
	TeamRepository ports.TeamRepository = nil
)

func Player(name string, tx *world.Tx) (*player.Player, bool) {
	for p := range tx.Players() {
		pl := p.(*player.Player)
		if strings.EqualFold(pl.Name(), name) {
			return pl, true
		}
	}
	return nil, false
}
