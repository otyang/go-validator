package validator

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

// WithFormTagsAsFieldNames configures the validator to use form tags as field names. e.g form:”name"
// It extracts field names from the "json" struct tags and ensures consistent
// validation behavior for both form data and JSON bodies.
func WithFormTagsAsFieldNames() ValidateOption {
	return func(v *Validator) {
		v.validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

// WithCustomValidationAlphaSpace returns an Options function that adds a custom validation
// named "alpha_space" to the validator. This validation checks if the field value
// contains only alphabetic characters and spaces.
//
// This can be useful for validating user input fields like names, addresses, or
// descriptions where only letters and spaces are allowed.
func WithCustomValidationAlphaSpace() ValidateOption {
	alphaSpace := func(fl validator.FieldLevel) bool {
		return regexp.MustCompile("^[a-zA-Z ]+$").MatchString(fl.Field().String())
	}

	return func(v *Validator) {
		v.validator.RegisterValidation("alpha_space", alphaSpace)
	}
}
