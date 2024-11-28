package message

import "github.com/df-mc/dragonfly/server/player"

type Writer struct{}

func (m Writer) Write(p *player.Player, s string) {
	p.Message(s)
}
