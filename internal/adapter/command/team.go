package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/restartfu/solar/internal/core/message"
	"github.com/restartfu/solar/internal/core/team"
	"github.com/restartfu/solar/internal/ports"
)

type TeamCreate struct {
	playerAllower
	db         ports.Database
	subscriber chat.Subscriber

	Sub  cmd.SubCommand `cmd:"create"`
	Name string
}

func NewTeam(sub chat.Subscriber, db ports.Database) cmd.Command {
	return cmd.New("team", "", nil,
		TeamCreate{subscriber: sub, db: db},
		TeamInvite{subscriber: sub, db: db},
	)
}

func (t TeamCreate) Run(src cmd.Source, out *cmd.Output, _ *world.Tx) {
	p := src.(*player.Player)

	_, ok := t.db.LoadTeam(t.Name)
	if ok {
		Messenger.Message(p, message.Team.CreateAlreadyExists(t.Name))
		return
	}

	tm := team.NewTeam(t.Name, p.Name())
	t.subscriber.Message(message.Team.CreateSuccess(t.Name, p.Name()))

	t.db.SaveTeam(tm)
}

type TeamInvite struct {
	playerAllower
	db         ports.Database
	subscriber chat.Subscriber

	Sub    cmd.SubCommand `cmd:"invite"`
	Target []cmd.Target
}

func (t TeamInvite) Run(src cmd.Source, out *cmd.Output, _ *world.Tx) {
	p := src.(*player.Player)
	target := t.Target[0].(*player.Player)

	tm, ok := t.db.LoadMemberTeam(p.Name())
	if !ok {
		Messenger.Message(p, message.Team.NotInTeam())
		return
	}

	if _, ok = t.db.LoadMemberTeam(target.Name()); ok {
		Messenger.Message(p, message.Team.TargetAlreadyInTeam(target.Name()))
		return
	}

	Messenger.Message(p, message.Team.InviteSent(target.Name()))
	Messenger.Message(target, message.Team.InviteReceived(tm.DisplayName))
}
