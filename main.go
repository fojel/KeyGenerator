package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	. "github.com/lxn/walk/declarative"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var app AppMainWindow
	app.Generator = &PasswordGeneratorValidator{}
	app.Validator = &PasswordGeneratorValidator{}

	if _, err := (MainWindow{
		AssignTo: &app.MainWindow,
		Title:    "Password Generator",
		Size:     Size{Width: 300, Height: 150},
		Layout:   VBox{},
		Children: []Widget{
			LineEdit{
				AssignTo: &app.PasswordLineEdit,
				ReadOnly: true,
			},
			PushButton{
				Text:      "Generate Password",
				OnClicked: app.GeneratePasswordClicked,
			},
		},
	}).Run(); err != nil {
		log.Fatalf("Error al ejecutar la ventana principal: %v", err)
	}
}

func (app *AppMainWindow) GeneratePasswordClicked() {
	length := 12
	password := app.Generator.Generate(length)

	if app.Validator.Validate(password) {
		app.PasswordLineEdit.SetText(password)
		return
	}
	fmt.Println("The generated password does not meet the length requirements.")
}
