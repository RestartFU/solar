package handler

import (
	"github.com/restartfu/solar/internal/core/domain/class"
	"sync/atomic"

	"github.com/df-mc/dragonfly/server/player"
)

var _ player.Handler = &PlayerHandler{}

type PlayerHandler struct {
	player.NopHandler
	activeClass class.Class
	activeArea  atomic.Value
}

func NewPlayerHandler(class class.Class) *PlayerHandler {
	h := &PlayerHandler{
		activeClass: class,
	}
	return h
}
