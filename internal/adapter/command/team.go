package command

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/restartfu/solar/internal/core/team"
	"github.com/restartfu/solar/internal/ports"
	"io"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/restartfu/solar/internal/core/message"
)

type TeamCreate struct {
	playerAllower
	db           ports.Database
	globalWriter io.StringWriter

	Sub  cmd.SubCommand `cmd:"create"`
	Name string
}

func NewTeam(w io.StringWriter, db ports.Database) cmd.Command {
	return cmd.New("team", "", nil,
		TeamCreate{globalWriter: w, db: db},
	)
}

func (t TeamCreate) Run(src cmd.Source, out *cmd.Output, _ *world.Tx) {
	p := src.(*player.Player)

	_, ok := t.db.LoadTeam(t.Name)
	if ok {
		Writer.Write(p, message.Team.CreateAlreadyExists(t.Name))
		return
	}

	tm := team.NewTeam(t.Name, p.Name())
	_, _ = t.globalWriter.WriteString(message.Team.CreateSuccess(t.Name, p.Name()))

	t.db.SaveTeam(tm)
}
