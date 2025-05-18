package utils

import (
	"errors"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
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

// ValidateFileSize securely checks if the file size is within allowed limits
func ValidateFileSize(file *multipart.FileHeader, maxSize int64) bool {
	return file != nil && file.Size > 0 && file.Size <= maxSize
}

// ValidateFileType securely checks file content type using both content sniffing and header validation
func ValidateFileType(file *multipart.FileHeader, allowedTypes []string) bool {
	if file == nil {
		return false
	}

	// Validate using content sniffing
	contentOk, err := validateContentType(file, allowedTypes)
	if err == nil && contentOk {
		return true
	}

	// Fallback to header validation
	return validateHeaderType(file, allowedTypes)
}

// ValidateFileExtension checks for allowed file extensions (case-insensitive)
func ValidateFileExtension(file *multipart.FileHeader, allowedExtensions []string) bool {
	if file == nil {
		return false
	}

	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(file.Filename), "."))
	normalizedExtensions := make([]string, len(allowedExtensions))
	for i, e := range allowedExtensions {
		normalizedExtensions[i] = strings.ToLower(strings.TrimPrefix(e, "."))
	}

	return slices.Contains(normalizedExtensions, ext)
}

// ValidateFileName checks for secure filenames and prevents path traversal
func ValidateFileName(file *multipart.FileHeader) bool {
	if file == nil {
		return false
	}

	cleaned := filepath.Clean(file.Filename)
	base := filepath.Base(cleaned)

	// Prevent empty filenames and directory traversal
	if cleaned == "." || cleaned == ".." || base != cleaned {
		return false
	}

	// Block special characters
	forbidden := []string{"<", ">", ":", "\"", "/", "\\", "|", "?", "*"}
	return !strings.ContainsAny(base, strings.Join(forbidden, ""))
}

// ValidateFileNameLength ensures filename length is within limits
func ValidateFileNameLength(file *multipart.FileHeader, maxLength int) bool {
	if file == nil {
		return false
	}
	return len(file.Filename) <= maxLength
}

// Helper functions
func validateContentType(file *multipart.FileHeader, allowedTypes []string) (bool, error) {
	f, err := file.Open()
	if err != nil {
		return false, err
	}
	defer f.Close()

	buffer := make([]byte, 512)
	n, err := f.Read(buffer)
	if err != nil && err != io.EOF {
		return false, err
	}

	detectedType := http.DetectContentType(buffer[:n])
	return slices.Contains(allowedTypes, detectedType), nil
}

func validateHeaderType(file *multipart.FileHeader, allowedTypes []string) bool {
	contentType := file.Header.Get("Content-Type")
	if contentType == "" {
		return false
	}

	// Parse media type to handle parameters like charset
	mimeType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return false
	}

	return slices.Contains(allowedTypes, mimeType)
}

// Example validation chain
func ValidateUploadedFile(file *multipart.FileHeader) error {
	if !ValidateFileName(file) {
		return errors.New("invalid filename")
	}

	if !ValidateFileNameLength(file, 255) {
		return errors.New("filename too long")
	}

	if !ValidateFileSize(file, (10 * 1024 * 1024)) { // 10MB limit
		return errors.New("file too large")
	}

	allowedTypes := []string{"image/jpeg", "image/png"}
	if !ValidateFileType(file, allowedTypes) {
		return errors.New("invalid file type")
	}

	allowedExts := []string{"jpg", "jpeg", "png"}
	if !ValidateFileExtension(file, allowedExts) {
		return errors.New("invalid file extension")
	}

	return nil
}
