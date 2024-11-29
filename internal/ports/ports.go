package ports

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/restartfu/solar/internal/core/domain"
)

type Database interface {
	LoadTeam(name string) (domain.Team, bool)
	LoadMemberTeam(name string) (domain.Team, bool)
	SaveTeam(team domain.Team)

	LoadUser(name string) (domain.User, bool)
	SaveUser(user domain.User)
}
type Messenger interface {
	Message(p *player.Player, s string)
}
