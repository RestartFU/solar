package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/restartfu/solar/internal/core"
)

type AllowerPlayer struct{}

func (AllowerPlayer) Allow(src cmd.Source) bool {
	_, ok := src.(*player.Player)
	return ok
}

type EnumUserInvitations string

func (u EnumUserInvitations) Type() string {
	return "user_invitations"
}

func (u EnumUserInvitations) Options(src cmd.Source) []string {
	p, ok := src.(*player.Player)
	if !ok {
		return nil
	}

	usr, ok := core.UserRepository.FindByName(p.Name())
	if !ok {
		return nil
	}

	return usr.Invitations.ActiveKeys()
}
