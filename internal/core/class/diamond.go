package class

import "github.com/df-mc/dragonfly/server/item"

type diamond struct {
	naked
}

func (diamond) Tiers() [4]item.ArmourTier {
	return [4]item.ArmourTier{
		item.ArmourTierDiamond{},
		item.ArmourTierDiamond{},
		item.ArmourTierDiamond{},
		item.ArmourTierDiamond{},
	}
}
