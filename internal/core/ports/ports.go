package ports

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/restartfu/solar/internal/core/domain/model"
)

type UserRepository interface {
	FindByName(name string) (model.User, bool)
	FindAll() model.User
	Save(model.User)
}

type TeamRepository interface {
	FindByMemberName(name string) (model.Team, bool)
	FindByName(name string) (model.Team, bool)
	FindAll() model.Team
	Save(model.Team)
}

type Messenger interface {
	Message(p *player.Player, s string)
}
