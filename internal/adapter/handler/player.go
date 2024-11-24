package handler

import (
	"sync/atomic"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/restartfu/solar/internal/core/class"
)

type PlayerHandler struct {
	player.NopHandler

	activeClass class.Class
	activeArea  atomic.Value
}

func NewPlayerHandler(p *player.Player) *PlayerHandler {
	h := &PlayerHandler{
		activeClass: class.Of(p),
	}
	return h
}
