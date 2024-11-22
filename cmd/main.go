package main

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/restartfu/gophig"
	"github.com/restartfu/solar/internal/adapter/command"
	"github.com/restartfu/solar/internal/adapter/handler"
	"log/slog"
	"os"
)

func main() {
	for _, c := range []cmd.Command{
		cmd.New("debug", "", nil,
			command.DebugActiveClass{},
		),
	} {
		cmd.Register(c)
	}

	chat.Global.Subscribe(chat.StdoutSubscriber{})
	conf, err := readConfig(slog.Default())
	if err != nil {
		panic(err)
	}

	srv := conf.New()
	srv.CloseOnProgramEnd()

	srv.Listen()
	for p := range srv.Accept() {
		playerHandler := handler.NewPlayerHandler(p)
		inventoryHandler := handler.NewInventoryHandler(playerHandler)
		armourHandler := handler.NewArmourHandler(playerHandler)

		p.Inventory().Handle(inventoryHandler)
		p.Armour().Handle(armourHandler)
		p.Handle(playerHandler)
	}
}

func readConfig(log *slog.Logger) (server.Config, error) {
	g := gophig.NewGophig[server.UserConfig]("./config.toml", gophig.TOMLMarshaler{}, 0777)

	c, err := g.LoadConf()
	if os.IsNotExist(err) {
		err = g.SaveConf(server.DefaultConfig())
		if err != nil {
			return server.Config{}, err
		}
	}
	return c.Config(log)
}
