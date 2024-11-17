package class

import (
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/inventory"
	"github.com/df-mc/dragonfly/server/player"
	"reflect"
	"time"
)

func all() []Class {
	return []Class{
		Diamond, Bard, Rogue, Archer,
	}
}

var (
	forever = time.Duration(1<<63 - 1)

	Naked   = Class{}
	Diamond = Class{
		armourTiers: [4]item.ArmourTier{
			item.ArmourTierDiamond{}, item.ArmourTierDiamond{}, item.ArmourTierDiamond{}, item.ArmourTierDiamond{},
		},
	}
	Bard = Class{
		armourTiers: [4]item.ArmourTier{
			item.ArmourTierGold{}, item.ArmourTierGold{}, item.ArmourTierGold{}, item.ArmourTierGold{},
		},
		effects: []effect.Effect{
			effect.New(effect.Speed{}, 2, forever),
		},
	}
	Rogue = Class{
		armourTiers: [4]item.ArmourTier{
			item.ArmourTierChain{}, item.ArmourTierChain{}, item.ArmourTierChain{}, item.ArmourTierChain{},
		},
		effects: []effect.Effect{
			effect.New(effect.Speed{}, 2, forever),
		},
	}
	Archer = Class{
		armourTiers: [4]item.ArmourTier{
			item.ArmourTierLeather{}, item.ArmourTierLeather{}, item.ArmourTierLeather{}, item.ArmourTierLeather{},
		},
		effects: []effect.Effect{
			effect.New(effect.Speed{}, 2, forever),
		},
	}
)

type Class struct {
	armourTiers [4]item.ArmourTier
	effects     []effect.Effect
}

func Is(initial, expected Class) bool {
	return reflect.DeepEqual(initial, expected)
}

func Of(p *player.Player) Class {
	for _, c := range all() {
		if armourIs(p.Armour(), c) {
			return c
		}
	}
	return Naked
}

func armourIs(armour *inventory.Armour, expected Class) bool {
	tiers := expected.armourTiers
	return pieceIsTier(armour.Helmet(), tiers[0]) &&
		pieceIsTier(armour.Chestplate(), tiers[1]) &&
		pieceIsTier(armour.Leggings(), tiers[2]) &&
		pieceIsTier(armour.Boots(), tiers[3])
}

func pieceIsTier(piece item.Stack, tier item.ArmourTier) bool {
	var pieceTier item.ArmourTier
	switch pce := piece.Item().(type) {
	case item.Helmet:
		pieceTier = pce.Tier
	case item.Chestplate:
		pieceTier = pce.Tier
	case item.Leggings:
		pieceTier = pce.Tier
	case item.Boots:
		pieceTier = pce.Tier
	}

	return pieceTier == tier
}
