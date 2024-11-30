package ports

import (
	"github.com/df-mc/dragonfly/server/player"
)

type Condition[T any] func(T) bool

type Identifiable interface {
	DisplayName() string
}

type Repository[T Identifiable] interface {
	Find(conds ...Condition[T]) (T, bool)
	Save(v T)
}

type Messenger interface {
	Message(p *player.Player, s string)
}
