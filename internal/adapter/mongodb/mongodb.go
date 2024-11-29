package mongodb

import "github.com/restartfu/solar/internal/core/team"

type DatabaseAdapter struct{}

func (d DatabaseAdapter) LoadTeam(name string) (team.Team, bool) {
	//TODO implement me
	panic("implement me")
}

func (d DatabaseAdapter) LoadMemberTeam(name string) (team.Team, bool) {
	//TODO implement me
	panic("implement me")
}

func (d DatabaseAdapter) SaveTeam(team team.Team) {
	//TODO implement me
	panic("implement me")
}
