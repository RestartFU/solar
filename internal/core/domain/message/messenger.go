package message

import "github.com/df-mc/dragonfly/server/player"

type Messenger struct{}

func (m Messenger) Message(p *player.Player, s string) {
	p.Message(s)
}
