package model

type Role int

const (
	Member Role = iota
	Admin
)

func (r Role) String() string {
	return []string{"MEMBER", "ADMIN"}[r]
}
