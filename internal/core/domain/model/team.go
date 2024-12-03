package model

import "strings"

type Importance int

const (
	ImportanceMinimal Importance = iota
	ImportancePartial
	ImportanceFull
)

type Team struct {
	Name        string
	DisplayName string

	Members []TeamMember
}

type TeamMember struct {
	DisplayName string
	Importance  Importance
}

func NewTeam(name string, leader string) Team {
	return Team{
		Name:        strings.ToLower(name),
		DisplayName: name,
		Members: []TeamMember{
			{
				DisplayName: leader,
				Importance:  ImportanceFull,
			},
		},
	}
}

func (t Team) FindMemberByNameAndImportance(name string, importance Importance) (TeamMember, bool) {
	for _, m := range t.Members {
		if strings.EqualFold(m.DisplayName, name) && m.Importance >= importance {
			return m, true
		}
	}
	return TeamMember{}, false
}
