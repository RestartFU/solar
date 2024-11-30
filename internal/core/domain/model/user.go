package model

import (
	"github.com/restartfu/solar/internal/core/domain/entity"
	"slices"
	"strings"
	"time"
)

const (
	invitationValidityLength = time.Hour
)

type User struct {
	identity    *entity.Identity
	invitations []*entity.Expirable
}

func NewUser(displayName string) User {
	return User{
		identity: &entity.Identity{
			Name:        strings.ToLower(displayName),
			DisplayName: displayName,
		},
	}
}

func (u User) DisplayName() string {
	return u.identity.DisplayName
}

func (u User) Invitations() (keys []string) {
	invitations := slices.Clone(u.invitations)

	for _, inv := range invitations {
		if !inv.Expired() {
			keys = append(keys, inv.Key)
		}
	}

	return keys
}

func (u User) WithInvitation(team string) (User, bool) {
	for _, inv := range u.Invitations() {
		if strings.EqualFold(team, inv) {
			return u, false
		}
	}
	u.invitations = append(u.invitations, &entity.Expirable{Key: team, Expiration: time.Now().Add(invitationValidityLength)})
	return u, true
}

func (u User) WithoutInvitations() User {
	u.invitations = nil
	return u
}
