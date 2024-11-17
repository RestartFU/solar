package handler

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"sync/atomic"

	"github.com/df-mc/dragonfly/server/item/inventory"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/restartfu/solar/internal/core/class"
)

type Handler struct {
	player.NopHandler
	inventory.Handler

	p           *player.Player
	activeClass atomic.Value
}

func NewHandler(p *player.Player) *Handler {
	h := &Handler{p: p}
	h.activeClass.Store(class.Of(p))
	return h
}

// HandleTake ...
func (h *Handler) HandleTake(_ *event.Context, slot int, _ item.Stack) {
	armour := h.p.Armour()

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
	h.p.Message(text.Colourf("<blue>your current class:</blue> <yellow>%s</yellow>", newClass.Name()))

	h.activeClass.Store(newClass)
}

// HandlePlace ...
func (h *Handler) HandlePlace(_ *event.Context, _ int, stack item.Stack) {
	armour := h.p.Armour()

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
	h.p.Message(text.Colourf("<blue>your current class:</blue> <yellow>%s</yellow>", newClass.Name()))

	h.activeClass.Store(newClass)
}

// HandleDrop ...
func (h *Handler) HandleDrop(_ *event.Context, slot int, _ item.Stack) {
	armour := h.p.Armour()

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
	h.p.Message(text.Colourf("<blue>your current class:</blue> <yellow>%s</yellow>", newClass.Name()))

	h.activeClass.Store(newClass)
}

func updateClass(p *player.Player, oldClass, newClass class.Class) {
	if class.Is(oldClass, newClass) {
		return
	}

	for _, eff := range oldClass.Effects() {
		p.RemoveEffect(eff.Type())
	}

	for _, eff := range newClass.Effects() {
		p.AddEffect(eff)
	}
}
