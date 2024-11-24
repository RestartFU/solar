package message

import (
	"github.com/restartfu/solar/internal/core/class"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type classMessages struct{}

func (classMessages) Enabled(c class.Class) string {
	return text.Colourf("<green>Enabled class</green><grey>:</grey> <yellow>%s</yellow>", class.NameOf(c))
}
