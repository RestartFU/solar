package model

import (
	"github.com/restartfu/solar/internal/core/domain/entity"
	"github.com/restartfu/solar/internal/ports"
	"slices"
	"strings"
)

type Team struct {
	identity *entity.Identity
	members  []*entity.ImportanceIdentity
}

func NewTeam(teamDisplayName string, leaderName string) Team {
	return Team{
		identity: &entity.Identity{
			Name:        strings.ToLower(teamDisplayName),
			DisplayName: teamDisplayName,
		},
		members: []*entity.ImportanceIdentity{
			{
				Identity: entity.Identity{
					Name:        strings.ToLower(leaderName),
					DisplayName: leaderName,
				},
				Importance: entity.ImportanceFull,
			},
		},
	}
}

func (t Team) DisplayName() string {
	return t.identity.DisplayName
}

func ConditionMemberImportance[T *entity.ImportanceIdentity](importance entity.Importance) func(v T) bool {
	return func(v T) bool {
		val := any(v).(*entity.ImportanceIdentity)
		return val.Importance >= importance
	}
}

func (t Team) Member(name string, cond ports.Condition[*entity.ImportanceIdentity]) (*entity.ImportanceIdentity, bool) {
	vals := slices.Clone(t.members)
	for _, val := range vals {
		if strings.EqualFold(val.Name, name) && cond(val) {
			return val, true
		}
	}
	return nil, false
}

func (t Team) Members(cond ports.Condition[*entity.ImportanceIdentity]) (members []*entity.ImportanceIdentity) {
	vals := slices.Clone(t.members)

	for _, val := range vals {
		if cond(val) {
			members = append(members, val)
		}
	}
	return members
}

func (t Team) WithMember(member *entity.ImportanceIdentity) Team {
	members := slices.Clone(t.members)
	members = append(members, member)

	t.members = members
	return t
}

func (t Team) WithoutMember(member *entity.ImportanceIdentity) Team {
	members := slices.Clone(t.members)
	i := slices.Index(members, member)
	if i == -1 {
		return t
	}
	slices.Delete(members, i, i)

	t.members = members
	return t
}
