package class

import (
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/inventory"
)

type naked struct {
	inventory.NopHandler
}

func (n naked) Tiers() [4]item.ArmourTier {
	return [4]item.ArmourTier{}
}

func (n naked) Effects() []effect.Type {
	return []effect.Type{}
}
