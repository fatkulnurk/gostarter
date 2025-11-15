package validation

import (
	"net"
	"strconv"
	"strings"
)

// parseTagToRules mengubah string tag validate jadi slice Rule.
//
// Contoh:
//
//	"required,minlen=3,maxlen=10,email"
//	"min=18,max=60"
func parseTagToRules(tag string) []Rule {
	parts := strings.Split(tag, ",")
	var rules []Rule

	for _, part := range parts {
		p := strings.TrimSpace(part)
		if p == "" {
			continue
		}

		switch {
		case p == RuleRequired:
			rules = append(rules, Required(""))

		case strings.HasPrefix(p, RuleStrMinLength):
			nStr := strings.TrimPrefix(p, RuleStrMinLength)
			n, err := strconv.Atoi(nStr)
			if err == nil {
				rules = append(rules, StrMinLength(n, ""))
			}

		case strings.HasPrefix(p, RuleStrMaxLength):
			nStr := strings.TrimPrefix(p, RuleStrMaxLength)
			n, err := strconv.Atoi(nStr)
			if err == nil {
				rules = append(rules, StrMaxLength(n, ""))
			}

		case strings.HasPrefix(p, RuleNumMin):
			nStr := strings.TrimPrefix(p, RuleNumMin)
			n, err := strconv.ParseFloat(nStr, 64)
			if err == nil {
				rules = append(rules, Min(n, ""))
			}

		case strings.HasPrefix(p, RuleNumMax):
			nStr := strings.TrimPrefix(p, RuleNumMax)
			n, err := strconv.ParseFloat(nStr, 64)
			if err == nil {
				rules = append(rules, Max(n, ""))
			}

		case p == RuleEmail:
			rules = append(rules, Custom(func(field string, value any) *Error {
				s, ok := value.(string)
				if !ok {
					return &Error{
						Field:   field,
						Message: "harus berupa string (email)",
					}
				}
				s = strings.TrimSpace(s)
				if s == "" {
					// Biar rule Required yang handle kalau dipakai
					return nil
				}
				if !strings.Contains(s, "@") {
					return &Error{
						Field:   field,
						Message: "format email tidak valid",
					}
				}
				return nil
			}))
		case p == "ip":
			rules = append(rules, Custom(func(field string, value any) *Error {
				s, ok := value.(string)
				if !ok {
					return &Error{
						Field:   field,
						Message: ErrorMessageString,
					}
				}
				s = strings.TrimSpace(s)
				if s == "" {
					// Biar rule Required yang handle kalau dipakai
					return nil
				}
				ip := net.ParseIP(s)
				if ip == nil {
					return &Error{
						Field:   field,
						Message: ErrorMessageInvalidIP,
					}
				}
				return nil
			}))
		case p == "ipv4":
			rules = append(rules, Custom(func(field string, value any) *Error {
				s, ok := value.(string)
				if !ok {
					return &Error{
						Field:   field,
						Message: ErrorMessageString,
					}
				}
				s = strings.TrimSpace(s)
				if s == "" {
					// Biar rule Required yang handle kalau dipakai
					return nil
				}
				ip := net.ParseIP(s)
				if ip == nil || ip.To4() == nil {
					return &Error{
						Field:   field,
						Message: ErrorMessageInvalidIPv4,
					}
				}
				return nil
			}))
		case p == "ipv6":
			rules = append(rules, Custom(func(field string, value any) *Error {
				s, ok := value.(string)
				if !ok {
					return &Error{
						Field:   field,
						Message: ErrorMessageString,
					}
				}
				s = strings.TrimSpace(s)
				if s == "" {
					// Biar rule Required yang handle kalau dipakai
					return nil
				}
				ip := net.ParseIP(s)
				if ip == nil || ip.To16() == nil {
					return &Error{
						Field:   field,
						Message: ErrorMessageInvalidIPv6,
					}
				}
				return nil
			}))

		default:
			// Tag tidak dikenal -> di-skip saja
			panic("tag tidak dikenal: " + p)
		}
	}

	return rules
}
