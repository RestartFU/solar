package repository

import (
	"github.com/restartfu/solar/internal/ports"
	"maps"
	"strings"
	"sync"
)

type Repository[T ports.Identifiable] struct {
	values map[string]T
	sync.Mutex
}

func NewUserRepository[T ports.Identifiable]() *Repository[T] {
	return &Repository[T]{
		values: make(map[string]T),
	}
}

func (r *Repository[T]) Find(conds ...ports.Condition[T]) (obj []T, foundAny bool) {
	r.Lock()
	vals := maps.Clone(r.values)
	r.Unlock()

	for _, v := range vals {
		for _, cond := range conds {
			if cond(v) {
				obj = append(obj, v)
				foundAny = true
			}
		}
	}

	return obj, false
}

func (r *Repository[T]) Save(u T) {
	r.Lock()
	r.values[strings.ToLower(u.DisplayName())] = u
	r.Unlock()
}
