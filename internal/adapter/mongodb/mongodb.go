package mongodb

import (
	"github.com/restartfu/solar/internal/core"
)

type DatabaseAdapter struct{}

func (d DatabaseAdapter) LoadTeam(name string) (core.Team, bool) {
	//TODO implement me
	panic("implement me")
}

func (d DatabaseAdapter) LoadMemberTeam(name string) (core.Team, bool) {
	//TODO implement me
	panic("implement me")
}

func (d DatabaseAdapter) SaveTeam(team core.Team) {
	//TODO implement me
	panic("implement me")
}
