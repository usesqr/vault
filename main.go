package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/usesqr/vault/ui/component"
	"github.com/usesqr/vault/ui/dialog"
)

func main() {
	a := app.New()
	w := a.NewWindow("Sqr Vault")
	w.Resize(fyne.Size{Width: 800, Height: 600})
	w.SetMainMenu(component.CreateMainMenu())

	hello := widget.NewLabel("Welcome to Sqr Vault.")
	w.SetContent(container.NewVBox(
		hello,
	))
	dialog.ShowFirstStartDialog(w, func(pass string, ok bool) {
		if !ok {
			os.Exit(1)
		}

		println(pass)
	})

	w.ShowAndRun()
}