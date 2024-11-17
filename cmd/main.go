package main

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/restartfu/solar/internal/adapter/command"
)

func main() {
	for _, c := range []cmd.Command{
		cmd.New("debug", "", nil,
			command.DebugActiveClass{},
		),
	} {
		cmd.Register(c)
	}
}
