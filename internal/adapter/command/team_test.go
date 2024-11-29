package command_test

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/restartfu/solar/internal/adapter/command"
	"github.com/restartfu/solar/internal/core"
	"github.com/restartfu/solar/internal/core/domain"
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
	mockTeam = domain.NewTeam(mockTeamName, mockPlayerName)
	mockUser = domain.NewUser(mockPlayerName)
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
				mockDatabase.EXPECT().LoadTeam(mockTeamName).Return(domain.Team{}, false)
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
				core.Subscriber = mockSubscriber
				mockMessenger := testutil.NewMessenger(t)
				core.Messenger = mockMessenger
				mockDatabaseAdapter := mocks.NewMockDatabase(ctrl)
				core.Database = mockDatabaseAdapter

				cmd.Register(command.NewTeam())

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
				mockDatabase.EXPECT().LoadUser(mockTargetPlayerName).Return(mockUser, true)
				mockDatabase.EXPECT().SaveUser(mockUser.WithInvitation(mockTeamName))

				mockMessenger.EXPECT(
					message.Team.InviteSent(mockTargetPlayerName),
					message.Team.InviteReceived(mockTeamName),
				)
			},
		},
		{
			name: "target data could not be loaded",
			setup: func(t *testing.T,
				mockSubscriber *testutil.Subscriber,
				mockMessenger *testutil.Messenger,
				mockDatabase *mocks.MockDatabase,
				mockPlayer *player.Player,
				tx *world.Tx,
			) {
				mockDatabase.EXPECT().LoadMemberTeam(mockPlayerName).Return(mockTeam, true)
				mockDatabase.EXPECT().LoadMemberTeam(mockTargetPlayerName).Return(mockTeam, false)
				mockDatabase.EXPECT().LoadUser(mockTargetPlayerName).Return(domain.User{}, false)
				mockMessenger.EXPECT(message.Error.LoadUserDataError(mockTargetPlayerName))
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
				mockDatabase.EXPECT().LoadMemberTeam(mockPlayerName).Return(domain.Team{}, false)
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
				core.Subscriber = mockSubscriber
				mockMessenger := testutil.NewMessenger(t)
				core.Messenger = mockMessenger
				mockDatabaseAdapter := mocks.NewMockDatabase(ctrl)
				core.Database = mockDatabaseAdapter

				cmd.Register(command.NewTeam())
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

func TestTeamJoin(t *testing.T) {
	mockTeam := mockTeam.WithoutMember(mockPlayerName)
	mockTeam = mockTeam.WithMember(mockTargetPlayerName, domain.RoleLeader)

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
			name: "player successfully joined team",
			setup: func(t *testing.T,
				mockSubscriber *testutil.Subscriber,
				mockMessenger *testutil.Messenger,
				mockDatabase *mocks.MockDatabase,
				mockPlayer *player.Player,
				tx *world.Tx,
			) {
				mockDatabase.EXPECT().LoadUser(mockPlayerName).Return(mockUser.WithInvitation(mockTeamName), true).Times(2)
				mockDatabase.EXPECT().LoadMemberTeam(mockPlayerName).Return(mockTeam, false)
				mockDatabase.EXPECT().SaveUser(mockUser)
				mockDatabase.EXPECT().SaveTeam(mockTeam.WithMember(mockPlayerName, domain.RoleMember))

				mockMessenger.EXPECT(
					message.Team.PlayerJoined(mockPlayerName),
					message.Team.PlayerJoined(mockPlayerName),
				)
			},
		},
		{
			name: "player is already in a team",
			setup: func(t *testing.T,
				mockSubscriber *testutil.Subscriber,
				mockMessenger *testutil.Messenger,
				mockDatabase *mocks.MockDatabase,
				mockPlayer *player.Player,
				tx *world.Tx,
			) {
				mockDatabase.EXPECT().LoadUser(mockPlayerName).Return(mockUser.WithInvitation(mockTeamName), true)
				mockDatabase.EXPECT().LoadMemberTeam(mockPlayerName).Return(mockTeam, true)
				mockMessenger.EXPECT(
					message.Team.AlreadyInTeam(),
				)
			},
		},
		{
			name: "could not load user data",
			setup: func(t *testing.T,
				mockSubscriber *testutil.Subscriber,
				mockMessenger *testutil.Messenger,
				mockDatabase *mocks.MockDatabase,
				mockPlayer *player.Player,
				tx *world.Tx,
			) {
				mockDatabase.EXPECT().LoadUser(mockPlayerName).Return(mockUser.WithInvitation(mockTeamName), true)
				mockDatabase.EXPECT().LoadMemberTeam(mockPlayerName).Return(domain.Team{}, false)
				mockDatabase.EXPECT().LoadUser(mockPlayerName).Return(domain.User{}, false)
				mockMessenger.EXPECT(
					message.Error.LoadUserDataError(mockPlayerName),
				)
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			testutil.MockWorld(func(tx *world.Tx) {
				ctrl := gomock.NewController(t)

				mockSubscriber := testutil.NewSubscriber(t)
				core.Subscriber = mockSubscriber
				mockMessenger := testutil.NewMessenger(t)
				core.Messenger = mockMessenger
				mockDatabaseAdapter := mocks.NewMockDatabase(ctrl)
				core.Database = mockDatabaseAdapter

				cmd.Register(command.NewTeam())
				_ = testutil.MockPlayer(tx, mockTargetPlayerName)
				mockPlayer := testutil.MockPlayer(tx, mockPlayerName)
				if tc.setup != nil {
					tc.setup(t, mockSubscriber, mockMessenger, mockDatabaseAdapter, mockPlayer, tx)
				}

				mockPlayer.ExecuteCommand("/team join " + mockTeamName)
			})
		})
	}
}

func TestTeamLeave(t *testing.T) {
	mockTeam := mockTeam.WithoutMember(mockPlayerName)
	mockTeam = mockTeam.WithMember(mockTargetPlayerName, domain.RoleLeader)
	mockTeam = mockTeam.WithMember(mockPlayerName, domain.RoleMember)

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
			name: "player successfully left team",
			setup: func(t *testing.T,
				mockSubscriber *testutil.Subscriber,
				mockMessenger *testutil.Messenger,
				mockDatabase *mocks.MockDatabase,
				mockPlayer *player.Player,
				tx *world.Tx,
			) {
				mockDatabase.EXPECT().LoadMemberTeam(mockPlayerName).Return(mockTeam, true)
				mockDatabase.EXPECT().SaveTeam(mockTeam.WithoutMember(mockPlayerName))

				mockMessenger.EXPECT(
					message.Team.PlayerLeft(mockPlayerName),
					message.Team.PlayerLeft(mockPlayerName),
				)
			},
		},
		{
			name: "player is not in a team",
			setup: func(t *testing.T,
				mockSubscriber *testutil.Subscriber,
				mockMessenger *testutil.Messenger,
				mockDatabase *mocks.MockDatabase,
				mockPlayer *player.Player,
				tx *world.Tx,
			) {
				mockDatabase.EXPECT().LoadMemberTeam(mockPlayerName).Return(domain.Team{}, false)
				mockMessenger.EXPECT(
					message.Team.NotInTeam(),
				)
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			testutil.MockWorld(func(tx *world.Tx) {
				ctrl := gomock.NewController(t)

				mockSubscriber := testutil.NewSubscriber(t)
				core.Subscriber = mockSubscriber
				mockMessenger := testutil.NewMessenger(t)
				core.Messenger = mockMessenger
				mockDatabaseAdapter := mocks.NewMockDatabase(ctrl)
				core.Database = mockDatabaseAdapter

				cmd.Register(command.NewTeam())
				_ = testutil.MockPlayer(tx, mockTargetPlayerName)
				mockPlayer := testutil.MockPlayer(tx, mockPlayerName)
				if tc.setup != nil {
					tc.setup(t, mockSubscriber, mockMessenger, mockDatabaseAdapter, mockPlayer, tx)
				}

				mockPlayer.ExecuteCommand("/team leave")
			})
		})
	}
}
