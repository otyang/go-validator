package validator

import (
	"errors"
	"fmt"

	en "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// Validator struct type contains a map of validation errors.
type Validator struct {
	Errors     map[string]string
	validator  *validator.Validate
	translator ut.Translator
}

type ValidateOption func(v *Validator)

// New creates a new Validator instance with optional configuration.
func New(opts ...ValidateOption) *Validator {
	v := validator.New()
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator(en.Locale())
	err := en_translations.RegisterDefaultTranslations(v, trans)
	if err != nil {
		panic(err)
	}

	validatorObj := &Validator{Errors: make(map[string]string), validator: v, translator: trans}
	for _, opt := range opts {
		opt(validatorObj)
	}

	return validatorObj
}

// Valid checks if there are any validation errors.
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// AddError adds a validation error for a given field.
func (v *Validator) AddError(key, message string) {
	if v.Errors == nil {
		v.Errors = make(map[string]string)
	}

	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Reset clears any existing validation errors.
func (v *Validator) Reset() {
	v.Errors = make(map[string]string)
}

// ValidateStruct validates a struct using the configured validator and translator.
func (v *Validator) ValidateStruct(vPtr any) error {
	if v.Errors == nil {
		v.Errors = make(map[string]string)
	}

	err := v.validator.Struct(vPtr)
	if err != nil {
		validatorErrs, ok := err.(validator.ValidationErrors)
		if ok {
			var errs []error

			for _, err := range validatorErrs {
				translatedErr := err.Translate(v.translator)
				v.AddError(err.Field(), translatedErr)
				errs = append(errs, fmt.Errorf(translatedErr))
			}

			var errMerged error
			for _, err := range errs {
				errMerged = errors.Join(errMerged, err)
			}

			return errMerged
		}

		return err
	}

	return nil
}
