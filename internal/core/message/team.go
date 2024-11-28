package message

import "github.com/sandertv/gophertunnel/minecraft/text"

type teamMessages struct{}

func (teamMessages) CreateSuccess(name string, leader string) string {
	return text.Colourf("<yellow>Team <blue>%s</blue> has been <green>created</green> by <grey>%s</grey></yellow>", name, leader)
}

func (teamMessages) CreateAlreadyExists(name string) string {
	return text.Colourf("<red>Team with name '%s' already exists</red>", name)
}
