package model

type Importance int

const (
	ImportanceMinimal Importance = iota
	ImportancePartial
	ImportanceFull
)

type TeamMember struct {
	DisplayName string
	Importance  Importance
}
