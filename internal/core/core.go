package core

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/restartfu/solar/internal/core/domain/model"
	"github.com/restartfu/solar/internal/core/message"
	"github.com/restartfu/solar/internal/ports"
	"strings"
)

var (
	Messenger      ports.Messenger              = message.Messenger{}
	Subscriber     chat.Subscriber              = chat.StdoutSubscriber{}
	UserRepository ports.Repository[model.User] = nil
	TeamRepository ports.Repository[model.Team] = nil
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
