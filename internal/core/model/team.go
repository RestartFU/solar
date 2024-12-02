package model

import "strings"

type Team struct {
	DisplayName string

	Members []TeamMember
}

func NewTeam(name string, leader string) Team {
	return Team{
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
