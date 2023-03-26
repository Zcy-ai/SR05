// Package gui build and launch GUI for this app
package gui

import (
	"fmt"
	"os"

	"sr05_ac4/handler"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// binding used to change the text of received message in the GUI
var (
	messageReceived binding.String
)

func LaunchGUI() {
	a := app.New()

	w := a.NewWindow("SR05")

	messageToSend := binding.NewString()
	messageReceived = binding.NewString()
	//speed := binding.NewFloat()

	entry := widget.NewEntry()
	entry.Bind(messageToSend)

	label := widget.NewLabel("")
	label.Bind(messageReceived)

	card := widget.NewCard("Received message", "", label)

	// slider := widget.NewSlider(0, 9999)
	// slider.SetValue(1000)
	// slider.Bind(speed)

	//create the handler with the interface to change received message, don't forget to close it at the end
	handler := handler.NewHandler()
	defer handler.Close()

	//send message once
	sendMessageButton := widget.NewButton("Send message", func() {
		text, err := messageToSend.Get()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Can't get message to send:", err)
			os.Exit(1)
		}

		handler.SendMessage(text)
	})

	form := widget.NewForm()
	form.Append("Message to send", entry)

	buttonContainer := container.New(layout.NewHBoxLayout(), sendMessageButton)
	mainContainer := container.New(layout.NewVBoxLayout(), form, buttonContainer, card)

	w.SetContent(mainContainer)

	handler.Run()

	w.ShowAndRun()
}
