package handler

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/restartfu/solar/internal/core/class"
	"sync/atomic"
)

type Handler struct {
	player.NopHandler
	activeClass *atomic.Value
}

func NewHandler(p *player.Player) *Handler {
	h := &Handler{}
	h.activeClass.Store(class.Of(p))
	return h
}
