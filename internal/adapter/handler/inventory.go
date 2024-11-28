package handler

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/inventory"
)

var _ inventory.Handler = &InventoryHandler{}

type InventoryHandler struct {
	playerHandler *PlayerHandler
}

func NewInventoryHandler(h *PlayerHandler) *InventoryHandler {
	return &InventoryHandler{playerHandler: h}
}

func (*InventoryHandler) HandleTake(*inventory.Context, int, item.Stack)  {}
func (*InventoryHandler) HandlePlace(*inventory.Context, int, item.Stack) {}
func (*InventoryHandler) HandleDrop(*inventory.Context, int, item.Stack)  {}
