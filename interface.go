package types

// Caster - интерфейс выражения типов
type (
	Caster interface {
		Cast(in, out interface{}) error
	}
)
