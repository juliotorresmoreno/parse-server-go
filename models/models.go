package models

type Error struct {
	error
	FieldName string
}

type Mode uint

const (
	ModeCreate Mode = iota
	ModeUpdate
)

type Model interface {
	Exists(id uint) bool
	FindById(id uint) (Model, error)
	Validate(Mode) []Error
	Insert() []Error
	Update(id uint) []Error
	Delete(id uint) error
}
