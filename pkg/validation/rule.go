package validation

import (
	"fmt"
	"strings"
)

// Rule adalah fungsi yang memvalidasi satu field.
type Rule func(field string, value any) *Error

// =========================
//  RULE BUILDER UMUM
// =========================

// Required : value tidak boleh kosong (nil atau string kosong / spasi).
// Jika message kosong (""), akan pakai default message.
func Required(message string) Rule {
	return func(field string, value any) *Error {
		msg := message
		if value == nil {
			if msg == "" {
				msg = ErrorMessageRequired
			}
			return &Error{Field: field, Message: msg}
		}

		// Jika string, cek kosong / spasi
		if s, ok := value.(string); ok {
			if strings.TrimSpace(s) == "" {
				if msg == "" {
					msg = ErrorMessageRequired
				}
				return &Error{Field: field, Message: msg}
			}
		}

		return nil
	}
}

// StrMinLength : panjang minimal untuk string.
func StrMinLength(min int, message string) Rule {
	return func(field string, value any) *Error {
		s, ok := value.(string)
		if !ok {
			return &Error{
				Field:   field,
				Message: ErrorMessageString,
			}
		}

		if len(s) < min {
			msg := message
			if msg == "" {
				msg = fmt.Sprintf(ErrorMessageMinLength, min)
			}
			return &Error{Field: field, Message: msg}
		}

		return nil
	}
}

// StrMaxLength : panjang maksimal untuk string.
func StrMaxLength(max int, message string) Rule {
	return func(field string, value any) *Error {
		s, ok := value.(string)
		if !ok {
			return &Error{
				Field:   field,
				Message: ErrorMessageString,
			}
		}

		if len(s) > max {
			msg := message
			if msg == "" {
				msg = fmt.Sprintf(ErrorMessageMaxLength, max)
			}
			return &Error{Field: field, Message: msg}
		}

		return nil
	}
}

// Min : nilai minimal untuk angka (int, uint, float32, float64, dll).
func Min(min float64, message string) Rule {
	return func(field string, value any) *Error {
		if value == nil {
			// Biarkan Required yang handle kalau dipakai
			return nil
		}

		num, ok := toFloat(value)
		if !ok {
			return &Error{
				Field:   field,
				Message: ErrorMessageNumber,
			}
		}

		if num < min {
			msg := message
			if msg == "" {
				msg = fmt.Sprintf(ErrorMessageMin, min)
			}
			return &Error{Field: field, Message: msg}
		}

		return nil
	}
}

// Max : nilai maksimal untuk angka (int, uint, float32, float64, dll).
func Max(max float64, message string) Rule {
	return func(field string, value any) *Error {
		if value == nil {
			return nil
		}

		num, ok := toFloat(value)
		if !ok {
			return &Error{
				Field:   field,
				Message: ErrorMessageNumber,
			}
		}

		if num > max {
			msg := message
			if msg == "" {
				msg = fmt.Sprintf(ErrorMessageMax, max)
			}
			return &Error{Field: field, Message: msg}
		}

		return nil
	}
}

// Custom : membuat rule kustom sendiri.
func Custom(f func(field string, value any) *Error) Rule {
	return func(field string, value any) *Error {
		return f(field, value)
	}
}

// helper: konversi berbagai tipe angka ke float64.
func toFloat(v any) (float64, bool) {
	switch n := v.(type) {
	case int:
		return float64(n), true
	case int8:
		return float64(n), true
	case int16:
		return float64(n), true
	case int32:
		return float64(n), true
	case int64:
		return float64(n), true
	case uint:
		return float64(n), true
	case uint8:
		return float64(n), true
	case uint16:
		return float64(n), true
	case uint32:
		return float64(n), true
	case uint64:
		return float64(n), true
	case float32:
		return float64(n), true
	case float64:
		return n, true
	default:
		return 0, false
	}
}
