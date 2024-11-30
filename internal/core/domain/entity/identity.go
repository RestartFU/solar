package entity

type ImportanceIdentity struct {
	Identity
	Importance
}

type Identity struct {
	Name        string
	DisplayName string
}
