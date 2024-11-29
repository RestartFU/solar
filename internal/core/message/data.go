package message

import "github.com/sandertv/gophertunnel/minecraft/text"

type errorMessages struct{}

func (errorMessages) LoadUserDataError(name string) string {
	return text.Colourf("<red>we could not load user data for '%s'; please contact a staff member</red>", name)
}
