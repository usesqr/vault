package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/usesqr/vault/ui/window"
)

func main() {
	a := app.New()
	window.CreateMainWindow(a).ShowAndRun()
}