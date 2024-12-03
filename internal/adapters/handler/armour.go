package handler

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/inventory"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/restartfu/solar/internal/core/domain/class"
	"github.com/restartfu/solar/internal/core/domain/message"
)

var _ inventory.Handler = &ArmourHandler{}

type ArmourHandler struct {
	playerHandler *PlayerHandler
}

func NewArmourHandler(h *PlayerHandler) *ArmourHandler {
	return &ArmourHandler{playerHandler: h}
}

// HandleTake ...
func (h *ArmourHandler) HandleTake(ctx *inventory.Context, slot int, _ item.Stack) {
	p := ctx.Val().(*player.Player)
	armour := p.Armour()

	tiers := [4]item.ArmourTier{}
	for _, s := range armour.Slots() {
		if !s.Empty() {
			switch pce := s.Item().(type) {
			case item.Helmet:
				tiers[0] = pce.Tier
			case item.Chestplate:
				tiers[1] = pce.Tier
			case item.Leggings:
				tiers[2] = pce.Tier
			case item.Boots:
				tiers[3] = pce.Tier
			}
		}
	}

	tiers[slot] = nil

	newClass := class.OfTiers(tiers)
	p.Message(message.Class.Enabled(newClass))

	h.playerHandler.activeClass = newClass
}

// HandlePlace ...
func (h *ArmourHandler) HandlePlace(ctx *inventory.Context, _ int, stack item.Stack) {
	p := ctx.Val().(*player.Player)
	armour := p.Armour()

	tiers := [4]item.ArmourTier{}
	for _, s := range armour.Slots() {
		if !s.Empty() {
			switch pce := s.Item().(type) {
			case item.Helmet:
				tiers[0] = pce.Tier
			case item.Chestplate:
				tiers[1] = pce.Tier
			case item.Leggings:
				tiers[2] = pce.Tier
			case item.Boots:
				tiers[3] = pce.Tier
			}
		}
	}

	switch pce := stack.Item().(type) {
	case item.Helmet:
		tiers[0] = pce.Tier
	case item.Chestplate:
		tiers[1] = pce.Tier
	case item.Leggings:
		tiers[2] = pce.Tier
	case item.Boots:
		tiers[3] = pce.Tier
	}

	newClass := class.OfTiers(tiers)
	p.Message(message.Class.Enabled(newClass))

	h.playerHandler.activeClass = newClass
}

// HandleDrop ...
func (h *ArmourHandler) HandleDrop(ctx *inventory.Context, slot int, _ item.Stack) {
	p := ctx.Val().(*player.Player)
	armour := p.Armour()

	tiers := [4]item.ArmourTier{}
	for _, s := range armour.Slots() {
		if !s.Empty() {
			switch pce := s.Item().(type) {
			case item.Helmet:
				tiers[0] = pce.Tier
			case item.Chestplate:
				tiers[1] = pce.Tier
			case item.Leggings:
				tiers[2] = pce.Tier
			case item.Boots:
				tiers[3] = pce.Tier
			}
		}
	}

	tiers[slot] = nil

	newClass := class.OfTiers(tiers)
	p.Message(message.Class.Enabled(newClass))

	h.playerHandler.activeClass = newClass
}
