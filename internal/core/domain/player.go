package domain

import (
	"strings"
)

type User struct {
	Name        string
	DisplayName string

	Invitations []string
}

func NewUser(name string) User {
	return User{
		Name:        strings.ToLower(name),
		DisplayName: name,
	}
}

func (u User) WithInvitation(from string) User {
	u.Invitations = append(u.Invitations, from)
	return u
}
