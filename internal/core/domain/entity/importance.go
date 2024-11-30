package entity

type Importance int

const (
	ImportanceMinimal Importance = iota
	ImportancePartial
	ImportanceFull
)
