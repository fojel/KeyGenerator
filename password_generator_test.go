package main

import (
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	pgv := &PasswordGeneratorValidator{}
	length := 12
	password := pgv.Generate(length)
	if len(password) != length {
		t.Errorf("Expected password length %d, but got %d", length, len(password))
	}
}

func TestValidatePassword(t *testing.T) {
	pgv := &PasswordGeneratorValidator{}
	validPassword := "ValidPass123"
	if !pgv.Validate(validPassword) {
		t.Errorf("Expected %s to be a valid password, but it was not", validPassword)
	}

	invalidPassword := "Weak"
	if pgv.Validate(invalidPassword) {
		t.Errorf("Expected %s to be an invalid password, but it was not", invalidPassword)
	}
}

func TestGeneratePasswordClicked(t *testing.T) {
	pga := &AppMainWindow{
		Generator: &PasswordGeneratorValidator{},
		Validator: &PasswordGeneratorValidator{},
	}

	pga.GeneratePasswordClicked()
	text := pga.PasswordLineEdit.Text()

	if text == "" {
		t.Error("Password line edit was not set after GeneratePasswordClicked")
	}

	if !pga.Validator.Validate(text) {
		t.Errorf("Generated password '%s' does not meet the length requirements", text)
	}
}
