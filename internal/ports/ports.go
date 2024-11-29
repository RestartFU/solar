package ports

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/restartfu/solar/internal/core"
)

type Database interface {
	LoadTeam(name string) (core.Team, bool)
	LoadMemberTeam(name string) (core.Team, bool)
	SaveTeam(team core.Team)
}
type Messenger interface {
	Message(p *player.Player, s string)
}
