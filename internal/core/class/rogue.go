package class

import "github.com/df-mc/dragonfly/server/item"

type rogue struct {
	naked
}

func (rogue) Tiers() [4]item.ArmourTier {
	return [4]item.ArmourTier{
		item.ArmourTierChain{},
		item.ArmourTierChain{},
		item.ArmourTierChain{},
		item.ArmourTierChain{},
	}
}
