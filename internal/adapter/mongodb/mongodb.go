package mongodb

import (
	"github.com/restartfu/solar/internal/core/domain"
)

type DatabaseAdapter struct{}

func (d DatabaseAdapter) LoadUser(name string) (domain.User, bool) {
	//TODO implement me
	panic("implement me")
}

func (d DatabaseAdapter) SaveUser(user domain.User) {
	//TODO implement me
	panic("implement me")
}

func (d DatabaseAdapter) LoadTeam(name string) (domain.Team, bool) {
	//TODO implement me
	panic("implement me")
}

func (d DatabaseAdapter) LoadMemberTeam(name string) (domain.Team, bool) {
	//TODO implement me
	panic("implement me")
}

func (d DatabaseAdapter) SaveTeam(team domain.Team) {
	//TODO implement me
	panic("implement me")
}
