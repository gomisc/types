package caster

import (
	"encoding/json"

	"git.eth4.dev/golibs/errors"

	"git.eth4.dev/golibs/types"
)

// ErrWrongCastTypes - ошибка типов параметров для выражения
const ErrWrongCastTypes = errors.Const("wrong types for cast. `in` must have save type with `out`, `out` must be interface")

// CastFunc - метод выражения типов
type CastFunc func(in, out interface{}) error

// Cast - реализация интерфейса выражения типов
func (cf CastFunc) Cast(in, out interface{}) error {
	return cf(in, out)
}

// DefaultCaster - метод выражения типов по умолчанию
// nolint: gocognit, gocyclo, cyclop
func Default() types.Caster {
	return CastFunc(func(in, out interface{}) error {
		var err error

		switch outv := out.(type) {
		case *int:
			if inv, ok := in.(int); ok {
				*outv = inv
			}
		case *int8:
			if inv, ok := in.(int8); ok {
				*outv = inv
			}
		case *int16:
			if inv, ok := in.(int16); ok {
				*outv = inv
			}
		case *int32:
			if inv, ok := in.(int32); ok {
				*outv = inv
			}
		case *int64:
			if inv, ok := in.(int64); ok {
				*outv = inv
			}
		case *uint:
			if inv, ok := in.(uint); ok {
				*outv = inv
			}
		case *uint8:
			if inv, ok := in.(uint8); ok {
				*outv = inv
			}
		case *uint16:
			if inv, ok := in.(uint16); ok {
				*outv = inv
			}
		case *uint32:
			if inv, ok := in.(uint32); ok {
				*outv = inv
			}
		case *uint64:
			if inv, ok := in.(uint64); ok {
				*outv = inv
			}
		case *float32:
			if inv, ok := in.(float32); ok {
				*outv = inv
			}
		case *float64:
			if inv, ok := in.(float64); ok {
				*outv = inv
			}
		case *uintptr:
			if inv, ok := in.(uintptr); ok {
				*outv = inv
			}
		case *string:
			if inv, ok := in.(string); ok {
				*outv = inv
			}

			if inv, ok := in.([]byte); ok {
				*outv = string(inv)
			}
		case *[]byte:
			if inv, ok := in.([]byte); ok {
				if err = json.Unmarshal(inv, out); err != nil {
					*outv = append(inv[:0:0], inv...) // nolint: gocritic
				}
			}
		default:
			if inv, ok := in.([]byte); ok {
				if err = json.Unmarshal(inv, out); err != nil {
					return errors.Wrap(err, "unmarshal casted value")
				}

				return nil
			}

			return ErrWrongCastTypes
		}

		return err
	})
}
