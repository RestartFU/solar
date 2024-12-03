package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/restartfu/solar/internal/core/domain/class"
	"github.com/restartfu/solar/internal/core/domain/message"
)

type DebugActiveClass struct {
	AllowerPlayer
	Sub cmd.SubCommand `cmd:"class"`
}

func (c DebugActiveClass) Run(src cmd.Source, out *cmd.Output, _ *world.Tx) {
	p := src.(*player.Player)
	out.Print(message.Class.Enabled(class.Of(p)))
}
