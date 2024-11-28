package testutil

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

func MockPlayer(tx *world.Tx, name string) *player.Player {
	opts := world.EntitySpawnOpts{}
	conf := player.Config{Name: name}
	return tx.AddEntity(opts.New(player.Type, conf)).(*player.Player)
}

func MockWorld(f func(tx *world.Tx)) {
	wrld := world.New()
	<-wrld.Exec(func(tx *world.Tx) {
		f(tx)
	})
}
