package handler

import "github.com/df-mc/dragonfly/server/item/inventory"

type InventoryHandler struct {
	inventory.NopHandler

	playerHandler *PlayerHandler
}

func NewInventoryHandler(h *PlayerHandler) *InventoryHandler {
	return &InventoryHandler{playerHandler: h}
}
