package handler

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"sync/atomic"

	"github.com/df-mc/dragonfly/server/item/inventory"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/restartfu/solar/internal/core/class"
)

type Handler struct {
	player.NopHandler
	inventory.Handler

	p           *player.Player
	activeClass *atomic.Value
}

func NewHandler(p *player.Player) *Handler {
	h := &Handler{p: p}
	h.activeClass.Store(class.Of(p))
	return h
}

// HandleTake ...
func (h *Handler) HandleTake(ctx *event.Context, slot int, stack item.Stack) {
	armour := h.p.Armour()

	tiers := [4]item.ArmourTier{
		armour.Helmet().Item().(item.Helmet).Tier,
		armour.Helmet().Item().(item.Chestplate).Tier,
		armour.Helmet().Item().(item.Leggings).Tier,
		armour.Helmet().Item().(item.Boots).Tier,
	}
	tiers[slot] = nil
	h.activeClass.Store(class.OfTiers(tiers))
}

// HandlePlace ...
func (h *Handler) HandlePlace(ctx *event.Context, slot int, stack item.Stack) {
	armour := h.p.Armour()

	tiers := [4]item.ArmourTier{
		armour.Helmet().Item().(item.Helmet).Tier,
		armour.Helmet().Item().(item.Chestplate).Tier,
		armour.Helmet().Item().(item.Leggings).Tier,
		armour.Helmet().Item().(item.Boots).Tier,
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

	h.activeClass.Store(class.OfTiers(tiers))
}

// HandleDrop ...
func (h *Handler) HandleDrop(ctx *event.Context, slot int, stack item.Stack) {
	armour := h.p.Armour()

	tiers := [4]item.ArmourTier{
		armour.Helmet().Item().(item.Helmet).Tier,
		armour.Helmet().Item().(item.Chestplate).Tier,
		armour.Helmet().Item().(item.Leggings).Tier,
		armour.Helmet().Item().(item.Boots).Tier,
	}
	tiers[slot] = nil
	h.activeClass.Store(class.OfTiers(tiers))
}
