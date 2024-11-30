package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/restartfu/solar/internal/core"
	"github.com/restartfu/solar/internal/core/domain/model"
	"github.com/restartfu/solar/internal/core/domain/repository"
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
	p, ok := src.(*player.Player)
	if !ok {
		return nil
	}

	usr, ok := core.UserRepository.Find(repository.ConditionName[model.User](p.Name()))
	if !ok {
		return nil
	}

	return usr.Invitations()
}
