package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type SignUpPayload struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,containsany=!@#?*"`
	Name     string `validate:"omitempty,min=4"`
}

func TestNew(t *testing.T) {
	validator := New()
	assert.NotNil(t, validator)
	assert.Empty(t, validator.Errors)
}

func TestValidator_ValidateStruct(t *testing.T) {
	ttt := SignUpPayload{
		Email:    "test@testcom", //   invalid email
		Password: "securepass",   //   missing required character
		Name:     "a12",          //   min length is not 4
	}

	v := New()

	err := v.ValidateStruct(ttt)
	assert.Error(t, err)

	assert.False(t, v.Valid())
}

func TestValidator_Valid(t *testing.T) {
	validator := New()
	assert.True(t, validator.Valid())

	validator.AddError("field1", "error1")
	assert.False(t, validator.Valid())
}

func TestValidator_AddError(t *testing.T) {
	validator := New()

	assert.True(t, validator.Valid()) // valid is true (no error)

	validator.AddError("field1", "error1")
	assert.Equal(t, map[string]string{"field1": "error1"}, validator.Errors)

	validator.AddError("field1", "error2") // Same key should not overwrite
	assert.Equal(t, map[string]string{"field1": "error1"}, validator.Errors)

	assert.False(t, validator.Valid())
}

func TestValidator_Reset(t *testing.T) {
	validator := New()
	validator.AddError("field1", "error1")
	validator.Reset()
	assert.Empty(t, validator.Errors)
}
