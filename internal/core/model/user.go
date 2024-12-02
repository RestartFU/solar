package model

import (
	"github.com/bedrock-gophers/cooldown/cooldown"
	"strings"
)

type User struct {
	Name        string
	DisplayName string
	Invitations cooldown.MappedCoolDown[string]
}

func NewUser(name string) User {
	return User{
		Name:        strings.ToLower(name),
		DisplayName: name,
		Invitations: cooldown.NewMappedCoolDown[string](),
	}
}
