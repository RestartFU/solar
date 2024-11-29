package domain

import (
	"maps"
	"strings"
)

const (
	RoleMember Role = iota
	RoleCaptain
	RoleLeader
)

type Role int

func (r Role) String() string {
	switch r {
	case RoleMember:
		return "Member"
	case RoleCaptain:
		return "Captain"
	case RoleLeader:
		return "Leader"
	}
	panic("should never happen")
}

type Team struct {
	Name        string
	DisplayName string
	Members     map[string]Role
}

func NewTeam(name string, leader string) Team {
	t := Team{
		Name:        strings.ToLower(name),
		DisplayName: name,
		Members: map[string]Role{
			leader: RoleLeader,
		},
	}
	return t
}

func (t Team) WithMember(name string, role Role) Team {
	t.Members = maps.Clone(t.Members)
	t.Members[name] = role
	return t
}
func (t Team) WithoutMember(name string) Team {
	t.Members = maps.Clone(t.Members)
	delete(t.Members, name)
	return t
}
