package entity

import "time"

type Expirable struct {
	Key        string
	Expiration time.Time
}

func (e Expirable) Expired() bool {
	return e.Expiration.Before(time.Now())
}
