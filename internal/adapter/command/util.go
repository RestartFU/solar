package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/restartfu/solar/internal/core"
)

type playerAllower struct{}

func (playerAllower) Allow(src cmd.Source) bool {
	_, ok := src.(*player.Player)
	return ok
}

type userInvitations string

func (u userInvitations) Type() string {
	return "user_invitations"
}

func (u userInvitations) Options(src cmd.Source) []string {
	p := src.(*player.Player)
	usr, ok := core.Database.LoadUser(p.Name())
	if !ok {
		return nil
	}

	return usr.Invitations
}
