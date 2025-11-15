package validation

import (
	"fmt"
	"strings"
)

// Rule adalah fungsi yang memvalidasi satu field.
type Rule func(field string, value any) *Error

// Custom : membuat rule kustom sendiri.
func Custom(f func(field string, value any) *Error) Rule {
	return func(field string, value any) *Error {
		return f(field, value)
	}
}

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

func Username(message string) Rule {
	return func(field string, value any) *Error {
		s, ok := value.(string)
		if !ok {
			return &Error{
				Field:   field,
				Message: ErrorMessageString,
			}
		}

		if len(s) < 6 || len(s) > 16 {
			msg := message
			if msg == "" {
				msg = fmt.Sprintf(ErrorMessageNumberBetween, float64(6), float64(16))
			}
			return &Error{Field: field, Message: msg}
		}

		for _, ch := range s {
			if !((ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') || ch == '_') {
				return &Error{
					Field:   field,
					Message: ErrorMessageInvalidUsername,
				}
			}
		}

		return nil
	}
}

func PhoneNumber(message string) Rule {
	return func(field string, value any) *Error {
		s, ok := value.(string)
		if !ok {
			return &Error{
				Field:   field,
				Message: ErrorMessageString,
			}
		}

		s = strings.TrimSpace(s)
		if s == "" {
			// Biarkan Required yang handle jika dipakai
			return nil
		}

		digitCount := 0
		for i, ch := range s {
			// hitung digit
			if ch >= '0' && ch <= '9' {
				digitCount++
				continue
			}

			// karakter non-digit yang diizinkan
			if ch == ' ' || ch == '-' || ch == '(' || ch == ')' {
				continue
			}
			if ch == '+' && i == 0 {
				continue
			}

			// selain itu dianggap tidak valid
			return &Error{
				Field:   field,
				Message: ErrorMessageInvalidPhone,
			}
		}

		if digitCount < 10 || digitCount > 15 {
			msg := message
			if msg == "" {
				msg = fmt.Sprintf(ErrorMessageNumberBetween, float64(10), float64(15))
			}
			return &Error{Field: field, Message: msg}
		}

		return nil
	}
}

func Password(message string) Rule {
	return func(field string, value any) *Error {
		s, ok := value.(string)
		if !ok {
			return &Error{
				Field:   field,
				Message: ErrorMessageString,
			}
		}

		s = strings.TrimSpace(s)
		if s == "" {
			// Biarkan Required yang handle jika dipakai
			return nil
		}

		if len(s) < 8 || len(s) > 16 {
			msg := message
			if msg == "" {
				msg = fmt.Sprintf(ErrorMessageNumberBetween, float64(8), float64(16))
			}
			return &Error{Field: field, Message: msg}
		}

		hasUpper := false
		hasLower := false
		hasDigit := false
		hasSpecial := false

		for _, ch := range s {
			switch {
			case ch >= 'A' && ch <= 'Z':
				hasUpper = true
			case ch >= 'a' && ch <= 'z':
				hasLower = true
			case ch >= '0' && ch <= '9':
				hasDigit = true
			default:
				hasSpecial = true
			}
		}

		if !(hasUpper && hasLower && hasDigit && hasSpecial) {
			msg := message
			if msg == "" {
				msg = ErrorMessageInvalidPassword
			}
			return &Error{Field: field, Message: msg}
		}

		return nil
	}
}
