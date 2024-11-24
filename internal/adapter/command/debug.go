package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/restartfu/solar/internal/core/class"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type DebugActiveClass struct {
	playerAllower
	Sub cmd.SubCommand `cmd:"class"`
}

func (c DebugActiveClass) Run(src cmd.Source, out *cmd.Output, _ *world.Tx) {
	p := src.(*player.Player)
	out.Print(text.Colourf("<blue>your current class:</blue> <yellow>%s</yellow>", class.NameOf(class.Of(p))))
}
