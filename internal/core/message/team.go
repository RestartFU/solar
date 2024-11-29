package message

import "github.com/sandertv/gophertunnel/minecraft/text"

type teamMessages struct{}

func (teamMessages) PlayerLeft(name string) string {
	return text.Colourf("<red>%s left the team</red>", name)
}

func (teamMessages) PlayerJoined(name string) string {
	return text.Colourf("<yellow>%s</yellow> <blue>joined the team</blue>", name)
}

func (teamMessages) InviteSent(targetName string) string {
	return text.Colourf("<green>Invite sent to <yellow>%s</yellow></green>", targetName)
}

func (teamMessages) InviteReceived(teamName string) string {
	return text.Colourf("<green><yellow>%s</yellow> invited you to their team</green>", teamName)
}

func (teamMessages) AlreadyInTeam() string {
	return text.Colourf("<red>You are already in a team</red>")
}

func (teamMessages) NotInTeam() string {
	return text.Colourf("<red>You are not in a team</red>")
}

func (teamMessages) TargetAlreadyInTeam(name string) string {
	return text.Colourf("<red>%s is already in a team</red>", name)
}

func (teamMessages) CreateSuccess(name string, leader string) string {
	return text.Colourf("<yellow>Team <blue>%s</blue> has been <green>created</green> by <grey>%s</grey></yellow>", name, leader)
}

func (teamMessages) CreateAlreadyExists(name string) string {
	return text.Colourf("<red>Team with name '%s' already exists</red>", name)
}
