package command_test

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/restartfu/solar/internal/adapter/command"
	"github.com/restartfu/solar/internal/core/message"
	"github.com/restartfu/solar/internal/core/team"
	"github.com/restartfu/solar/mocks"
	"github.com/restartfu/solar/pkg/testutil"
	"go.uber.org/mock/gomock"
	"strings"
	"testing"
)

const (
	mockPlayerName = "testPlayer"
	mockTeamName   = "testTeam"
)

func TestTeamCreate(t *testing.T) {
	mockTeam := team.NewTeam(mockTeamName, mockPlayerName)

	for _, tc := range []struct {
		name      string
		arguments []string
		setup     func(t *testing.T,
			mockStringWriter *testutil.StringWriter,
			mockMessageWriter *testutil.StringWriter,
			mockDatabase *mocks.MockDatabase,
			mockPlayer *player.Player,
			tx *world.Tx,
		)
	}{
		{
			name:      "team is successfully created",
			arguments: []string{mockTeamName},
			setup: func(t *testing.T,
				mockStringWriter *testutil.StringWriter,
				mockMessageWriter *testutil.StringWriter,
				mockDatabase *mocks.MockDatabase,
				mockPlayer *player.Player,
				tx *world.Tx,
			) {
				mockDatabase.EXPECT().LoadTeam(mockTeamName).Return(team.Team{}, false)
				mockStringWriter.EXPECT(message.Team.CreateSuccess(mockTeamName, mockPlayerName))
				mockDatabase.EXPECT().SaveTeam(mockTeam)
			},
		},
		{
			name:      "team with name already exists",
			arguments: []string{mockTeamName},
			setup: func(t *testing.T,
				mockStringWriter *testutil.StringWriter,
				mockMessageWriter *testutil.StringWriter,
				mockDatabase *mocks.MockDatabase,
				mockPlayer *player.Player,
				tx *world.Tx,
			) {
				mockDatabase.EXPECT().LoadTeam(mockTeamName).Return(team.Team{}, true)
				mockMessageWriter.EXPECT(message.Team.CreateAlreadyExists(mockTeamName))
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			testutil.MockWorld(func(tx *world.Tx) {
				ctrl := gomock.NewController(t)

				mockStringWriter := testutil.NewStringWriter(t)

				mockMessageWriter := testutil.NewStringWriter(t)
				command.Writer = mockMessageWriter

				mockDatabaseAdapter := mocks.NewMockDatabase(ctrl)
				cmd.Register(command.NewTeam(mockStringWriter, mockDatabaseAdapter))

				mockPlayer := testutil.MockPlayer(tx, mockPlayerName)
				if tc.setup != nil {
					tc.setup(t, mockStringWriter, mockMessageWriter, mockDatabaseAdapter, mockPlayer, tx)
				}

				mockPlayer.ExecuteCommand("/team create " + strings.Join(tc.arguments, " "))
			})
		})
	}
}
