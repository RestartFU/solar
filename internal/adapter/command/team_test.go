package command_test

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/restartfu/solar/internal/adapter/command"
	"github.com/restartfu/solar/internal/core"
	"github.com/restartfu/solar/internal/core/message"
	"github.com/restartfu/solar/mocks"
	"github.com/restartfu/solar/pkg/testutil"
	"go.uber.org/mock/gomock"
	"testing"
)

const (
	mockPlayerName       = "testPlayer"
	mockTargetPlayerName = "testTargetPlayer"
	mockTeamName         = "testTeam"
)

var (
	mockTeam = core.NewTeam(mockTeamName, mockPlayerName)
)

func TestTeamCreate(t *testing.T) {

	for _, tc := range []struct {
		name  string
		setup func(t *testing.T,
			mockSubscriber *testutil.Subscriber,
			mockMessenger *testutil.Messenger,
			mockDatabase *mocks.MockDatabase,
			mockPlayer *player.Player,
			tx *world.Tx,
		)
	}{
		{
			name: "team is successfully created",
			setup: func(t *testing.T,
				mockSubscriber *testutil.Subscriber,
				mockMessenger *testutil.Messenger,
				mockDatabase *mocks.MockDatabase,
				mockPlayer *player.Player,
				tx *world.Tx,
			) {
				mockDatabase.EXPECT().LoadTeam(mockTeamName).Return(core.Team{}, false)
				mockSubscriber.EXPECT(message.Team.CreateSuccess(mockTeamName, mockPlayerName))
				mockDatabase.EXPECT().SaveTeam(mockTeam)
			},
		},
		{
			name: "team with name already exists",
			setup: func(t *testing.T,
				mockSubscriber *testutil.Subscriber,
				mockMessenger *testutil.Messenger,
				mockDatabase *mocks.MockDatabase,
				mockPlayer *player.Player,
				tx *world.Tx,
			) {
				mockDatabase.EXPECT().LoadTeam(mockTeamName).Return(mockTeam, true)
				mockMessenger.EXPECT(message.Team.CreateAlreadyExists(mockTeamName))
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			testutil.MockWorld(func(tx *world.Tx) {
				ctrl := gomock.NewController(t)

				mockSubscriber := testutil.NewSubscriber(t)
				command.Subscriber = mockSubscriber
				mockMessenger := testutil.NewMessenger(t)
				command.Messenger = mockMessenger

				mockDatabaseAdapter := mocks.NewMockDatabase(ctrl)
				cmd.Register(command.NewTeam(mockDatabaseAdapter))

				mockPlayer := testutil.MockPlayer(tx, mockPlayerName)
				if tc.setup != nil {
					tc.setup(t, mockSubscriber, mockMessenger, mockDatabaseAdapter, mockPlayer, tx)
				}

				mockPlayer.ExecuteCommand("/team create " + mockTeamName)
			})
		})
	}
}

func TestTeamInvite(t *testing.T) {
	for _, tc := range []struct {
		name  string
		setup func(t *testing.T,
			mockSubscriber *testutil.Subscriber,
			mockMessenger *testutil.Messenger,
			mockDatabase *mocks.MockDatabase,
			mockPlayer *player.Player,
			tx *world.Tx,
		)
	}{
		{
			name: "target invited successfully",
			setup: func(t *testing.T,
				mockSubscriber *testutil.Subscriber,
				mockMessenger *testutil.Messenger,
				mockDatabase *mocks.MockDatabase,
				mockPlayer *player.Player,
				tx *world.Tx,
			) {
				mockDatabase.EXPECT().LoadMemberTeam(mockPlayerName).Return(mockTeam, true)
				mockDatabase.EXPECT().LoadMemberTeam(mockTargetPlayerName).Return(mockTeam, false)
				mockMessenger.EXPECT(
					message.Team.InviteSent(mockTargetPlayerName),
					message.Team.InviteReceived(mockTeamName))
			},
		},
		{
			name: "source not in a team",
			setup: func(t *testing.T,
				mockSubscriber *testutil.Subscriber,
				mockMessenger *testutil.Messenger,
				mockDatabase *mocks.MockDatabase,
				mockPlayer *player.Player,
				tx *world.Tx,
			) {
				mockDatabase.EXPECT().LoadMemberTeam(mockPlayerName).Return(core.Team{}, false)
				mockMessenger.EXPECT(message.Team.NotInTeam())
			},
		},
		{
			name: "target is already in a team",
			setup: func(t *testing.T,
				mockSubscriber *testutil.Subscriber,
				mockMessenger *testutil.Messenger,
				mockDatabase *mocks.MockDatabase,
				mockPlayer *player.Player,
				tx *world.Tx,
			) {
				mockDatabase.EXPECT().LoadMemberTeam(mockPlayerName).Return(mockTeam, true)
				mockDatabase.EXPECT().LoadMemberTeam(mockTargetPlayerName).Return(mockTeam, true)
				mockMessenger.EXPECT(message.Team.TargetAlreadyInTeam(mockTargetPlayerName))
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			testutil.MockWorld(func(tx *world.Tx) {
				ctrl := gomock.NewController(t)

				mockSubscriber := testutil.NewSubscriber(t)
				mockMessenger := testutil.NewMessenger(t)
				command.Messenger = mockMessenger

				mockDatabaseAdapter := mocks.NewMockDatabase(ctrl)
				cmd.Register(command.NewTeam(mockDatabaseAdapter))

				_ = testutil.MockPlayer(tx, mockTargetPlayerName)
				mockPlayer := testutil.MockPlayer(tx, mockPlayerName)
				if tc.setup != nil {
					tc.setup(t, mockSubscriber, mockMessenger, mockDatabaseAdapter, mockPlayer, tx)
				}

				mockPlayer.ExecuteCommand("/team invite " + mockTargetPlayerName)
			})
		})
	}
}
