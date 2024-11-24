package class

import "github.com/df-mc/dragonfly/server/item"

type archer struct {
	naked
}

func (archer) Tiers() [4]item.ArmourTier {
	return [4]item.ArmourTier{
		item.ArmourTierLeather{},
		item.ArmourTierLeather{},
		item.ArmourTierLeather{},
		item.ArmourTierLeather{},
	}
}
