package class

import (
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/inventory"
	"github.com/df-mc/dragonfly/server/player"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"reflect"
	"time"
)

type Class interface {
	Tiers() [4]item.ArmourTier
	Effects() []effect.Type
}

func All() []Class {
	return []Class{
		Diamond, Bard, Rogue, Archer,
	}
}

var (
	forever    = time.Duration(1<<63 - 1)
	titleCaser = cases.Title(language.English)

	Naked   = naked{}
	Diamond = diamond{}
	Bard    = bard{}
	Rogue   = rogue{}
	Archer  = archer{}
)

func Is(initial, expected Class) bool {
	return reflect.DeepEqual(initial, expected)
}

func Of(p *player.Player) Class {
	for _, c := range All() {
		if armourIs(p.Armour(), c) {
			return c
		}
	}
	return Naked
}

func OfTiers(tiers [4]item.ArmourTier) Class {
	for _, c := range All() {
		if reflect.DeepEqual(c.Tiers(), tiers) {
			return c
		}
	}
	return Naked
}

func NameOf(c Class) string {
	className := reflect.TypeOf(c).Name()
	return titleCaser.String(className)
}

func armourIs(armour *inventory.Armour, expected Class) bool {
	tiers := expected.Tiers()
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
