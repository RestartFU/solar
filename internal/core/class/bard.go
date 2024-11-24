package class

import "github.com/df-mc/dragonfly/server/item"

type bard struct {
	naked
}

func (bard) Tiers() [4]item.ArmourTier {
	return [4]item.ArmourTier{
		item.ArmourTierGold{},
		item.ArmourTierGold{},
		item.ArmourTierGold{},
		item.ArmourTierGold{},
	}
}
