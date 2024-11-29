package ports

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/restartfu/solar/internal/core/team"
)

type Database interface {
	LoadTeam(name string) (team.Team, bool)
	LoadMemberTeam(name string) (team.Team, bool)
	SaveTeam(team team.Team)
}
type MessageWriter interface {
	Write(p *player.Player, s string)
}
