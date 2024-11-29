package core

import "strings"

type Player struct {
	Name        string
	DisplayName string

	Invitations []string
}

func NewPlayer(name string) Player {
	return Player{
		Name:        strings.ToLower(name),
		DisplayName: name,
	}
}
