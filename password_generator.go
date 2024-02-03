package main

import (
	"github.com/lxn/walk"
	"math/rand"
)

type IPasswordGenerator interface {
	Generate(length int) string
}

type IPasswordValidator interface {
	Validate(password string) bool
}

type PasswordGeneratorValidator struct{}

type AppMainWindow struct {
	MainWindow       *walk.MainWindow
	PasswordLineEdit *walk.LineEdit
	Generator        IPasswordGenerator
	Validator        IPasswordValidator
}

func (pgv *PasswordGeneratorValidator) Generate(length int) string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+="
	password := make([]byte, length)
	for i := range password {
		password[i] = characters[rand.Intn(len(characters))]
	}
	return string(password)
}

func (pgv *PasswordGeneratorValidator) Validate(password string) bool {
	return len(password) >= 8
}
