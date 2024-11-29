package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/restartfu/solar/internal/core"
	"github.com/restartfu/solar/internal/core/domain"
	"github.com/restartfu/solar/internal/core/message"
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

	_, ok := core.Database.LoadTeam(t.Name)
	if ok {
		core.Messenger.Message(p, message.Team.CreateAlreadyExists(t.Name))
		return
	}

	tm := domain.NewTeam(t.Name, p.Name())
	core.Subscriber.Message(message.Team.CreateSuccess(t.Name, p.Name()))

	core.Database.SaveTeam(tm)
}

type TeamInvite struct {
	playerAllower

	Sub    cmd.SubCommand `cmd:"invite"`
	Target []cmd.Target
}

func (t TeamInvite) Run(src cmd.Source, out *cmd.Output, _ *world.Tx) {
	p := src.(*player.Player)
	target := t.Target[0].(*player.Player)

	tm, ok := core.Database.LoadMemberTeam(p.Name())
	if !ok {
		core.Messenger.Message(p, message.Team.NotInTeam())
		return
	}

	if _, ok = core.Database.LoadMemberTeam(target.Name()); ok {
		core.Messenger.Message(p, message.Team.TargetAlreadyInTeam(target.Name()))
		return
	}

	u, ok := core.Database.LoadUser(target.Name())
	if !ok {
		core.Messenger.Message(p, message.Error.LoadUserDataError(target.Name()))
		return
	}
	core.Database.SaveUser(u.WithInvitation(tm.DisplayName))

	core.Messenger.Message(p, message.Team.InviteSent(target.Name()))
	core.Messenger.Message(target, message.Team.InviteReceived(tm.DisplayName))
}

type TeamJoin struct {
	playerAllower

	Sub        cmd.SubCommand `cmd:"join"`
	Invitation userInvitations
}

func (t TeamJoin) Run(src cmd.Source, out *cmd.Output, tx *world.Tx) {
	p := src.(*player.Player)

	tm, ok := core.Database.LoadMemberTeam(p.Name())
	if ok {
		core.Messenger.Message(p, message.Team.AlreadyInTeam())
		return
	}

	u, ok := core.Database.LoadUser(p.Name())
	if !ok {
		core.Messenger.Message(p, message.Error.LoadUserDataError(p.Name()))
		return
	}

	u.Invitations = nil
	core.Database.SaveUser(u)
	tm = tm.WithMember(p.Name(), domain.RoleMember)
	core.Database.SaveTeam(tm)

	for name := range tm.Members {
		pl, online := core.Player(name, tx)
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

	tm, ok := core.Database.LoadMemberTeam(p.Name())
	if !ok {
		core.Messenger.Message(p, message.Team.NotInTeam())
		return
	}

	for name := range tm.Members {
		pl, online := core.Player(name, tx)
		if !online {
			continue
		}
		core.Messenger.Message(pl, message.Team.PlayerLeft(p.Name()))
	}

	core.Database.SaveTeam(tm.WithoutMember(p.Name()))
}
