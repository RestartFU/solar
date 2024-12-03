package model_test

import (
	"github.com/restartfu/solar/internal/core/domain/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewTeam(t *testing.T) {
	mockTeam := model.Team{
		Name:        "test",
		DisplayName: "TEST",
		Members: []model.TeamMember{
			{
				DisplayName: "test",
				Importance:  model.ImportanceFull,
			},
		},
	}

	tm := model.NewTeam("TEST", "test")
	require.Equal(t, mockTeam, tm)
}

func TestTeam_FindMemberByNameAndImportance(t *testing.T) {
	tm := model.NewTeam("TEST", "test")
	mockMember := model.TeamMember{
		DisplayName: "test2",
		Importance:  model.ImportanceMinimal,
	}
	tm.Members = append(tm.Members, mockMember)

	for _, tc := range []struct {
		name string

		memberName string
		expected   bool
		importance model.Importance
	}{
		{
			name:       "leader is found",
			memberName: "test",
			expected:   true,
			importance: model.ImportanceFull,
		},
		{
			name:       "member is not found",
			memberName: "test2",
			expected:   false,
			importance: model.ImportancePartial,
		},
		{
			name:       "member is found",
			memberName: "test2",
			expected:   true,
			importance: model.ImportanceMinimal,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, ok := tm.FindMemberByNameAndImportance(tc.memberName, tc.importance)
			require.Equal(t, tc.expected, ok)
		})
	}
}
