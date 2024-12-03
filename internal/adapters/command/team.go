package command

import (
	"github.com/restartfu/solar/internal/core/domain/message"
	"github.com/restartfu/solar/internal/core/domain/model"
	"time"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/restartfu/solar/internal/core"
)

const (
	teamInvitationDuration = time.Hour
)

type TeamCreate struct {
	AllowerPlayer

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

	_, ok := core.TeamRepository.FindByName(t.Name)
	if ok {
		core.Messenger.Message(p, message.Team.CreateAlreadyExists(t.Name))
		return
	}

	tm := model.NewTeam(t.Name, p.Name())
	core.Subscriber.Message(message.Team.CreateSuccess(t.Name, p.Name()))

	core.TeamRepository.Save(tm)
}

type TeamInvite struct {
	AllowerPlayer

	Sub    cmd.SubCommand `cmd:"invite"`
	Target []cmd.Target
}

func (t TeamInvite) Run(src cmd.Source, out *cmd.Output, _ *world.Tx) {
	p := src.(*player.Player)
	target := t.Target[0].(*player.Player)

	tm, ok := core.TeamRepository.FindByMemberName(p.Name())
	if !ok {
		core.Messenger.Message(p, message.Team.NotInTeam())
		return
	}

	if _, canInvite := tm.FindMemberByNameAndImportance(p.Name(), model.ImportancePartial); !canInvite {
		core.Messenger.Message(p, "TODO")
	}

	if _, ok = core.TeamRepository.FindByMemberName(target.Name()); ok {
		core.Messenger.Message(p, message.Team.TargetAlreadyInTeam(target.Name()))
		return
	}

	u, ok := core.UserRepository.FindByName(target.Name())
	if !ok {
		core.Messenger.Message(p, message.Error.LoadUserDataError(target.Name()))
		return
	}

	if u.Invitations.Key(tm.DisplayName).Active() {
		// already invited
		return
	}

	u.Invitations.Set(tm.DisplayName, teamInvitationDuration)
	core.UserRepository.Save(u)

	core.Messenger.Message(p, message.Team.InviteSent(target.Name()))
	core.Messenger.Message(target, message.Team.InviteReceived(tm.DisplayName))
}

type TeamJoin struct {
	AllowerPlayer

	Sub        cmd.SubCommand `cmd:"join"`
	Invitation EnumUserInvitations
}

func (t TeamJoin) Run(src cmd.Source, out *cmd.Output, tx *world.Tx) {
	p := src.(*player.Player)

	tm, ok := core.TeamRepository.FindByMemberName(p.Name())
	if ok {
		core.Messenger.Message(p, message.Team.AlreadyInTeam())
		return
	}

	u, ok := core.UserRepository.FindByName(p.Name())
	if !ok {
		core.Messenger.Message(p, message.Error.LoadUserDataError(p.Name()))
		return
	}

	clear(u.Invitations)
	core.UserRepository.Save(u)
	tm.Members = append(tm.Members, model.TeamMember{
		DisplayName: p.Name(),
		Importance:  model.ImportanceMinimal,
	})
	core.TeamRepository.Save(tm)

	for _, name := range tm.Members {
		pl, online := core.Player(name.DisplayName, tx)
		if !online {
			continue
		}
		core.Messenger.Message(pl, message.Team.PlayerJoined(p.Name()))
	}
}

type TeamLeave struct {
	AllowerPlayer

	Sub cmd.SubCommand `cmd:"leave"`
}

func (t TeamLeave) Run(src cmd.Source, out *cmd.Output, tx *world.Tx) {
	p := src.(*player.Player)

	tm, ok := core.TeamRepository.FindByMemberName(p.Name())
	if !ok {
		core.Messenger.Message(p, message.Team.NotInTeam())
		return
	}

	member, ok := tm.FindMemberByNameAndImportance(p.Name(), model.ImportanceMinimal)
	if !ok {
		core.Messenger.Message(p, message.Team.NotInTeam())
		return
	}

	if member.Importance == model.ImportanceFull {
		core.Messenger.Message(p, message.Team.LeaderLeave())
		return
	}

	for _, name := range tm.Members {
		pl, online := core.Player(name.DisplayName, tx)
		if !online {
			continue
		}
		core.Messenger.Message(pl, message.Team.PlayerLeft(p.Name()))
	}

	clear(tm.Members)
	core.TeamRepository.Save(tm)
}
