package utils

import (
	"regexp"
	"slices"
)

var (
	DatabaseVarcharMaxLength = 255
	DatabaseTextMaxLength    = 65535
	DatabaseTinyintMaxValue  = 127
	DatabaseSmallintMaxValue = 32767
	DatabaseIntMaxValue      = 2147483647
	DatabaseBigintMaxValue   = 9223372036854775807
	DatabaseFloatMaxDigits   = 7
	DatabaseDoubleMaxDigits  = 15
	DatabaseDecimalMaxDigits = 65
)

var (
	ValidationErrorRequired            = "%s is required"
	ValidationErrorInvalidEmail        = "%s is not a valid email address"
	ValidationErrorInvalidPhone        = "%s is not a valid phone number"
	ValidationErrorInvalidPassword     = "%s must contain at least 8 characters including uppercase, lowercase, number, and special character"
	ValidationErrorInvalidUsername     = "%s must be 3-16 characters and can only contain letters, numbers, and underscores"
	ValidationErrorInvalidURL          = "%s is not a valid URL"
	ValidationErrorInvalidIPv4         = "%s is not a valid IPv4 address"
	ValidationErrorInvalidDate         = "%s is not a valid date format (YYYY-MM-DD)"
	ValidationErrorInvalidAlphaNumeric = "%s must contain only letters and numbers"
	ValidationErrorInvalidUUID         = "%s is not a valid UUID"
	ValidationErrorInvalidJSON         = "%s is not a valid JSON"
	ValidationErrorInvalidHexColor     = "%s is not a valid hex color code"
	ValidationErrorInvalidCreditCard   = "%s is not a valid credit card number"
	ValidationErrorInvalidPostalCode   = "%s is not a valid postal code"
	ValidationErrorInvalidBase64       = "%s is not a valid base64 string"
	ValidationErrorOneOf               = "%s must be one of %s"
	ValidationErrorLowercase           = "%s must be lowercase"
	ValidationErrorUppercase           = "%s must be uppercase"
	ValidationErrorNumber              = "%s must be a number"
	ValidationErrorFloat               = "%s must be a float"
	ValidationErrorInt                 = "%s must be an integer"
	ValidationErrorMin                 = "%s must be greater than or equal to %d"
	ValidationErrorMax                 = "%s must be less than or equal to %d"
	ValidationErrorMinLength           = "%s must be at least %d characters long"
	ValidationErrorMaxLength           = "%s must be at most %d characters long"
	ValidationErrorGreaterThan         = "%s must be greater than %d"
	ValidationErrorLessThan            = "%s must be less than %d"
	ValidationErrorGreaterThanEqual    = "%s must be greater than or equal to %d"
	ValidationErrorLessThanEqual       = "%s must be less than or equal to %d"
	ValidationErrorGreaterThanEqualStr = "%s must be greater than or equal to %s"
	ValidationErrorLessThanEqualStr    = "%s must be less than or equal to %s"
	ValidationErrorNotEqual            = "%s must not be equal to %s"
	ValidationErrorNotEqualStr         = "%s must not be equal to %s"
	ValidationErrorNotIn               = "%s must not be in %s"
)

func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func ValidatePhone(phone string) bool {
	re := regexp.MustCompile(`^(\+[0-9]{1,3})?[0-9]{10,14}$`)
	return re.MatchString(phone)
}

func ValidatePassword(password string) bool {
	re := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`)
	return re.MatchString(password)
}

func ValidateUsername(username string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]{3,16}$`)
	return re.MatchString(username)
}

func ValidateURL(url string) bool {
	re := regexp.MustCompile(`^(http|https)://[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,}(/\S*)?$`)
	return re.MatchString(url)
}

func ValidateIPv4(ip string) bool {
	re := regexp.MustCompile(`^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$`)
	return re.MatchString(ip)
}

func ValidateDate(date string) bool {
	re := regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$`)
	return re.MatchString(date)
}

func ValidateAlphaNumeric(str string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return re.MatchString(str)
}

func OneOf(value string, options []string) bool {
	return slices.Contains(options, value)
}
