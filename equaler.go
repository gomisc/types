package types

import (
	"bytes"
	"strings"
)

// Equaler - интерфейс сравнения типов на равенство
type Equaler interface {
	Equal(x, Y interface{}) bool
}

// EqualFunc - метод сравнения типов на равенство
type EqualFunc func(x, y interface{}) bool

// Equal - реализация интерфейсу сравнения типов на равенство
func (lf EqualFunc) Equal(x, y interface{}) bool {
	return lf(x, y)
}

// PrimEqualer - метод сравнения значений примитивных типов
// nolint: gocognit, gocyclo, cyclop
func PrimEqualer() Equaler {
	return EqualFunc(func(x, y interface{}) bool {
		switch yv := y.(type) {
		case int:
			if xv, ok := x.(int); ok {
				return xv == yv
			}
		case int8:
			if xv, ok := x.(int8); ok {
				return xv == yv
			}
		case int16:
			if xv, ok := x.(int16); ok {
				return xv == yv
			}
		case int32:
			if xv, ok := x.(int32); ok {
				return xv == yv
			}
		case int64:
			if xv, ok := x.(int64); ok {
				return xv == yv
			}
		case uint:
			if xv, ok := x.(uint); ok {
				return xv == yv
			}
		case uint8:
			if xv, ok := x.(uint8); ok {
				return xv == yv
			}
		case uint16:
			if xv, ok := x.(uint16); ok {
				return xv == yv
			}
		case uint32:
			if xv, ok := x.(uint32); ok {
				return xv == yv
			}
		case uint64:
			if xv, ok := x.(uint64); ok {
				return xv == yv
			}
		case float32:
			if xv, ok := x.(float32); ok {
				return xv == yv
			}
		case float64:
			if xv, ok := x.(float64); ok {
				return xv == yv
			}
		case uintptr:
			if xv, ok := x.(uintptr); ok {
				return xv == yv
			}
		case string:
			if xv, ok := x.(string); ok {
				return strings.EqualFold(xv, yv)
			}
		case []byte:
			if xv, ok := x.([]byte); ok {
				return bytes.EqualFold(xv, yv)
			}
		case bool:
			if xv, ok := x.(bool); ok {
				return xv == yv
			}
		}

		return false
	})
}
