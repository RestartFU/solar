package ports

import (
	"github.com/df-mc/dragonfly/server/player"
	model2 "github.com/restartfu/solar/internal/core/domain/model"
)

type UserRepository interface {
	FindByName(name string) (model2.User, bool)
	FindAll() model2.User
	Save(model2.User)
}

type TeamRepository interface {
	FindByMemberName(name string) (model2.Team, bool)
	FindByName(name string) (model2.Team, bool)
	FindAll() model2.Team
	Save(model2.Team)
}

type Messenger interface {
	Message(p *player.Player, s string)
}
