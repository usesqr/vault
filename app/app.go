package app

import (
	"fyne.io/fyne/v2/app"
	"github.com/usesqr/vault/ui/window"
)

func Run() {
	a := app.New()
	window.CreateMainWindow(a).ShowAndRun()
}