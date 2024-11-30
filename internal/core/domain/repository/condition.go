package repository

import (
	"github.com/restartfu/solar/internal/core/domain/entity"
	"github.com/restartfu/solar/internal/core/domain/model"
	"strings"
)

func ConditionName[T model.Team | model.User](name string) func(v T) bool {
	return func(v T) bool {
		if val, ok := any(v).(interface{ DisplayName() string }); ok {
			return strings.EqualFold(val.DisplayName(), name)
		}
		return false
	}
}

func ConditionUserInTeam[T model.Team](username string) func(v T) bool {
	return func(v T) bool {
		val := model.Team(v)
		for _, member := range val.Members(model.ConditionMemberImportance(entity.ImportanceMinimal)) {
			if strings.EqualFold(member.Name, username) {
				return true
			}
		}
		return false
	}
}
