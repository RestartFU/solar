package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/restartfu/solar/internal/core"
	"github.com/restartfu/solar/internal/core/domain/entity"
	"github.com/restartfu/solar/internal/core/domain/model"
	"github.com/restartfu/solar/internal/core/domain/repository"
	"github.com/restartfu/solar/internal/core/message"
	"strings"
)

type TeamCreate struct {
	playerAllower

	Sub  cmd.SubCommand `cmd:"create"`
	Name string
}

func NewTeam() cmd.Command {
	return cmd.New("team", "", nil,
		TeamCreate{},
		TeamInvite{},
		TeamJoin{},
		TeamLeave{},
	)
}

func (t TeamCreate) Run(src cmd.Source, out *cmd.Output, _ *world.Tx) {
	p := src.(*player.Player)

	_, ok := core.TeamRepository.Find(repository.ConditionName[model.Team](t.Name))
	if ok {
		core.Messenger.Message(p, message.Team.CreateAlreadyExists(t.Name))
		return
	}

	tm := model.NewTeam(t.Name, p.Name())
	core.Subscriber.Message(message.Team.CreateSuccess(t.Name, p.Name()))

	core.TeamRepository.Save(tm)
}

type TeamInvite struct {
	playerAllower

	Sub    cmd.SubCommand `cmd:"invite"`
	Target []cmd.Target
}

func (t TeamInvite) Run(src cmd.Source, out *cmd.Output, _ *world.Tx) {
	p := src.(*player.Player)
	target := t.Target[0].(*player.Player)

	tm, ok := core.TeamRepository.Find(repository.ConditionUserInTeam(p.Name()))
	if !ok {
		core.Messenger.Message(p, message.Team.NotInTeam())
		return
	}

	if _, canInvite := tm.Member(p.Name(), model.ConditionMemberImportance(entity.ImportancePartial)); !canInvite {
		core.Messenger.Message(p, "TODO")
	}

	if _, ok = core.TeamRepository.Find(repository.ConditionUserInTeam(target.Name())); ok {
		core.Messenger.Message(p, message.Team.TargetAlreadyInTeam(target.Name()))
		return
	}

	u, ok := core.UserRepository.Find(repository.ConditionName[model.User](target.Name()))
	if !ok {
		core.Messenger.Message(p, message.Error.LoadUserDataError(target.Name()))
		return
	}

	u, couldInvite := u.WithInvitation(tm.DisplayName())
	if !couldInvite {
		// already invited
		return
	}
	core.UserRepository.Save(u)

	core.Messenger.Message(p, message.Team.InviteSent(target.Name()))
	core.Messenger.Message(target, message.Team.InviteReceived(tm.DisplayName()))
}

type TeamJoin struct {
	playerAllower

	Sub        cmd.SubCommand `cmd:"join"`
	Invitation userInvitations
}

func (t TeamJoin) Run(src cmd.Source, out *cmd.Output, tx *world.Tx) {
	p := src.(*player.Player)

	tm, ok := core.TeamRepository.Find(repository.ConditionUserInTeam(p.Name()))
	if ok {
		core.Messenger.Message(p, message.Team.AlreadyInTeam())
		return
	}

	u, ok := core.UserRepository.Find(repository.ConditionName[model.User](p.Name()))
	if !ok {
		core.Messenger.Message(p, message.Error.LoadUserDataError(p.Name()))
		return
	}

	core.UserRepository.Save(u.WithoutInvitations())
	tm = tm.WithMember(&entity.ImportanceIdentity{
		Identity: entity.Identity{
			Name:        strings.ToLower(p.Name()),
			DisplayName: p.Name(),
		},
		Importance: entity.ImportanceMinimal,
	})
	core.TeamRepository.Save(tm)

	for _, name := range tm.Members(model.ConditionMemberImportance(entity.ImportanceMinimal)) {
		pl, online := core.Player(name.Name, tx)
		if !online {
			continue
		}
		core.Messenger.Message(pl, message.Team.PlayerJoined(p.Name()))
	}
}

type TeamLeave struct {
	playerAllower

	Sub cmd.SubCommand `cmd:"leave"`
}

func (t TeamLeave) Run(src cmd.Source, out *cmd.Output, tx *world.Tx) {
	p := src.(*player.Player)

	tm, ok := core.TeamRepository.Find(repository.ConditionUserInTeam(p.Name()))
	if !ok {
		core.Messenger.Message(p, message.Team.NotInTeam())
		return
	}
	identity, ok := tm.Member(p.Name(), model.ConditionMemberImportance(entity.ImportanceMinimal))
	if !ok {
		core.Messenger.Message(p, message.Team.NotInTeam())
		return
	}

	if identity.Importance == entity.ImportanceFull {
		// leader cant leave
		return
	}

	for _, name := range tm.Members(model.ConditionMemberImportance(entity.ImportanceMinimal)) {
		pl, online := core.Player(name.Name, tx)
		if !online {
			continue
		}
		core.Messenger.Message(pl, message.Team.PlayerLeft(p.Name()))
	}

	core.TeamRepository.Save(tm.WithoutMember(identity))
}
