package auth

import (
	"goyave.dev/goyave/v3/validation"
	"unicode"
)

// Goyave provides a powerful, yet easy way to validate all incoming data, no matter
// its type or its format, thanks to a large number of validation rules.

// Incoming requests are validated using rules set, which associate rules
// with each expected field in the request.

// Learn more about validation here: https://goyave.dev/guide/basics/validation.html

// This is the validation rules for the "/echo" route, which is simply
// writing the input as a response.

func isValidPassWord(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 6 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

var (
	RegisterRequest = validation.RuleSet{
		"email":     {"required", "string", "between:3,125", "email"},
		"lastname":  {"required", "string", "between:3,125"},
		"firstname": {"required", "string", "between:3,125"},
		"password":  {"required", "string", "between:6,64"},
	}
)
