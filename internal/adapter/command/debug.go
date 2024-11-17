package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/restartfu/solar/internal/core/class"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type DebugActiveClass struct {
	playerAllower
	cmd.SubCommand `cmd:"active_class"`
}

func (DebugActiveClass) Run(src cmd.Source, out *cmd.Output) {
	p := src.(*player.Player)
	out.Print(text.Colourf("<blue>your current class:</blue> <yellow>%s</yellow>", class.Of(p).Name()))
}
