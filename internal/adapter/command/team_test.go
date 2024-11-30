package command_test

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/restartfu/solar/internal/adapter/command"
	"github.com/restartfu/solar/internal/core"
	"github.com/restartfu/solar/internal/core/domain/entity"
	"github.com/restartfu/solar/internal/core/domain/model"
	"github.com/restartfu/solar/internal/core/message"
	"github.com/restartfu/solar/mocks"
	"github.com/restartfu/solar/pkg/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"strings"
	"testing"
)

const (
	mockPlayerName       = "testPlayer"
	mockTargetPlayerName = "testTargetPlayer"
	mockTeamName         = "testTeam"
)

func TestTeamCreate(t *testing.T) {
	var (
		mockTeam = model.NewTeam(mockTeamName, mockPlayerName)
	)

	for _, tc := range []struct {
		name  string
		setup func(t *testing.T,
			mockSubscriber *testutil.Subscriber,
			mockMessenger *testutil.Messenger,
			mockTeamRepository *mocks.MockRepository[model.Team],
		)
	}{
		{
			name: "team is successfully created",
			setup: func(t *testing.T,
				mockSubscriber *testutil.Subscriber,
				mockMessenger *testutil.Messenger,
				mockTeamRepository *mocks.MockRepository[model.Team],
			) {
				mockTeamRepository.EXPECT().Find(gomock.Any()).Return(model.Team{}, false)
				mockSubscriber.EXPECT(message.Team.CreateSuccess(mockTeamName, mockPlayerName))
				mockTeamRepository.EXPECT().Save(mockTeam)
			},
		},
		{
			name: "team with name already exists",
			setup: func(t *testing.T,
				mockSubscriber *testutil.Subscriber,
				mockMessenger *testutil.Messenger,
				mockTeamRepository *mocks.MockRepository[model.Team],
			) {
				mockTeamRepository.EXPECT().Find(gomock.Any()).Return(mockTeam, true)
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
				mockRepositoryAdapter := mocks.NewMockRepository[model.Team](ctrl)
				core.TeamRepository = mockRepositoryAdapter

				cmd.Register(command.NewTeam())

				mockPlayer := testutil.MockPlayer(tx, mockPlayerName)
				if tc.setup != nil {
					tc.setup(t, mockSubscriber, mockMessenger, mockRepositoryAdapter)
				}

				mockPlayer.ExecuteCommand("/team create " + mockTeamName)
			})
		})
	}
}

func TestTeamInvite(t *testing.T) {
	var (
		mockTeam = model.NewTeam(mockTeamName, mockPlayerName)
	)

	for _, tc := range []struct {
		name  string
		setup func(t *testing.T,
			mockSubscriber *testutil.Subscriber,
			mockMessenger *testutil.Messenger,
			mockTeamRepository *mocks.MockRepository[model.Team],
			mockUserRepository *mocks.MockRepository[model.User],
		)
	}{
		{
			name: "target invited successfully",
			setup: func(t *testing.T,
				mockSubscriber *testutil.Subscriber,
				mockMessenger *testutil.Messenger,
				mockTeamRepository *mocks.MockRepository[model.Team],
				mockUserRepository *mocks.MockRepository[model.User],
			) {
				mockTargetUser := model.NewUser(mockTargetPlayerName)

				mockTeamRepository.EXPECT().Find(gomock.Any()).Return(mockTeam, true)
				mockTeamRepository.EXPECT().Find(gomock.Any()).Return(model.Team{}, false)
				mockUserRepository.EXPECT().Find(gomock.Any()).Return(mockTargetUser, true)

				mockUserRepository.EXPECT().Save(gomock.Any())
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
				mockTeamRepository *mocks.MockRepository[model.Team],
				mockUserRepository *mocks.MockRepository[model.User],
			) {
				mockTeamRepository.EXPECT().Find(gomock.Any()).Return(mockTeam, true)
				mockTeamRepository.EXPECT().Find(gomock.Any()).Return(mockTeam, false)
				mockUserRepository.EXPECT().Find(gomock.Any()).Return(model.User{}, false)
				mockMessenger.EXPECT(message.Error.LoadUserDataError(mockTargetPlayerName))
			},
		},
		{
			name: "source not in a team",
			setup: func(t *testing.T,
				mockSubscriber *testutil.Subscriber,
				mockMessenger *testutil.Messenger,
				mockTeamRepository *mocks.MockRepository[model.Team],
				_ *mocks.MockRepository[model.User],
			) {
				mockTeamRepository.EXPECT().Find(gomock.Any()).Return(model.Team{}, false)
				mockMessenger.EXPECT(message.Team.NotInTeam())
			},
		},
		{
			name: "target is already in a team",
			setup: func(t *testing.T,
				mockSubscriber *testutil.Subscriber,
				mockMessenger *testutil.Messenger,
				mockTeamRepository *mocks.MockRepository[model.Team],
				_ *mocks.MockRepository[model.User],
			) {
				mockTeamRepository.EXPECT().Find(gomock.Any()).Return(mockTeam, true)
				mockTeamRepository.EXPECT().Find(gomock.Any()).Return(mockTeam, true)
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
				mockTeamRepositoryAdapter := mocks.NewMockRepository[model.Team](ctrl)
				core.TeamRepository = mockTeamRepositoryAdapter
				mockUserRepositoryAdapter := mocks.NewMockRepository[model.User](ctrl)
				core.UserRepository = mockUserRepositoryAdapter

				cmd.Register(command.NewTeam())
				_ = testutil.MockPlayer(tx, mockTargetPlayerName)
				mockPlayer := testutil.MockPlayer(tx, mockPlayerName)
				if tc.setup != nil {
					tc.setup(t, mockSubscriber, mockMessenger, mockTeamRepositoryAdapter, mockUserRepositoryAdapter)
				}

				mockPlayer.ExecuteCommand("/team invite " + mockTargetPlayerName)
			})
		})
	}
}

func TestTeamJoin(t *testing.T) {
	for _, tc := range []struct {
		name  string
		setup func(t *testing.T,
			mockSubscriber *testutil.Subscriber,
			mockMessenger *testutil.Messenger,
			mockTeamRepository *mocks.MockRepository[model.Team],
			mockUserRepository *mocks.MockRepository[model.User],
		)
	}{
		{
			name: "player successfully joined team",
			setup: func(t *testing.T,
				mockSubscriber *testutil.Subscriber,
				mockMessenger *testutil.Messenger,
				mockTeamRepository *mocks.MockRepository[model.Team],
				mockUserRepository *mocks.MockRepository[model.User],
			) {
				var (
					mockTeam = model.NewTeam(mockTeamName, mockTargetPlayerName)
					mockUser = model.NewUser(mockPlayerName)
				)
				mockUser, ok := mockUser.WithInvitation(mockTeamName)
				require.True(t, ok)
				require.Len(t, mockUser.Invitations(), 1)

				mockTeamRepository.EXPECT().Find(gomock.Any()).Return(mockTeam, false)
				mockUserRepository.EXPECT().Find(gomock.Any()).Return(mockUser, true).Times(2)

				mockUserRepository.EXPECT().Save(mockUser.WithoutInvitations())
				mockTeamRepository.EXPECT().Save(
					mockTeam.WithMember(&entity.ImportanceIdentity{
						Identity: entity.Identity{
							Name:        strings.ToLower(mockPlayerName),
							DisplayName: mockPlayerName,
						},
						Importance: entity.ImportanceMinimal,
					}),
				)

				mockMessenger.EXPECT(
					message.Team.PlayerJoined(mockPlayerName),
					message.Team.PlayerJoined(mockPlayerName),
				)
			},
		},
		/*{
			name: "player is already in a team",
			setup: func(t *testing.T,
				mockSubscriber *testutil.Subscriber,
				mockMessenger *testutil.Messenger,
				mockTeamRepository *mocks.MockRepository[model.Team],
				mockUserRepository *mocks.MockRepository[model.User],
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
				mockTeamRepository *mocks.MockRepository[model.Team],
				mockUserRepository *mocks.MockRepository[model.User],
			) {
				mockDatabase.EXPECT().LoadUser(mockPlayerName).Return(mockUser.WithInvitation(mockTeamName), true)
				mockDatabase.EXPECT().LoadMemberTeam(mockPlayerName).Return(domain.Team{}, false)
				mockDatabase.EXPECT().LoadUser(mockPlayerName).Return(domain.User{}, false)
				mockMessenger.EXPECT(
					message.Error.LoadUserDataError(mockPlayerName),
				)
			},
		},*/
	} {
		t.Run(tc.name, func(t *testing.T) {
			testutil.MockWorld(func(tx *world.Tx) {
				ctrl := gomock.NewController(t)

				mockSubscriber := testutil.NewSubscriber(t)
				core.Subscriber = mockSubscriber
				mockMessenger := testutil.NewMessenger(t)
				core.Messenger = mockMessenger
				mockTeamRepositoryAdapter := mocks.NewMockRepository[model.Team](ctrl)
				core.TeamRepository = mockTeamRepositoryAdapter
				mockUserRepositoryAdapter := mocks.NewMockRepository[model.User](ctrl)
				core.UserRepository = mockUserRepositoryAdapter

				cmd.Register(command.NewTeam())
				_ = testutil.MockPlayer(tx, mockTargetPlayerName)
				mockPlayer := testutil.MockPlayer(tx, mockPlayerName)
				if tc.setup != nil {
					tc.setup(t, mockSubscriber, mockMessenger, mockTeamRepositoryAdapter, mockUserRepositoryAdapter)
				}

				mockPlayer.ExecuteCommand("/team join " + mockTeamName)
			})
		})
	}
}

/*func TestTeamLeave(t *testing.T) {
	mockTeam := mockTeam.WithoutMember(mockPlayerName)
	mockTeam = mockTeam.WithMember(mockTargetPlayerName, domain.RoleLeader)
	mockTeam = mockTeam.WithMember(mockPlayerName, domain.RoleMember)

	for _, tc := range []struct {
		name  string
		setup func(t *testing.T,
			mockSubscriber *testutil.Subscriber,
			mockMessenger *testutil.Messenger,
			mockDatabase *mocks.MockDatabase,
		)
	}{
		{
			name: "player successfully left team",
			setup: func(t *testing.T,
				mockSubscriber *testutil.Subscriber,
				mockMessenger *testutil.Messenger,
				mockDatabase *mocks.MockDatabase,
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
					tc.setup(t, mockSubscriber, mockMessenger, mockDatabaseAdapter)
				}

				mockPlayer.ExecuteCommand("/team leave")
			})
		})
	}
}*/
