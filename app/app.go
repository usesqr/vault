package app

import (
	"fyne.io/fyne/v2/app"
	"github.com/usesqr/vault/theme"
	"github.com/usesqr/vault/ui/window"
)

func Run() {
	a := app.New()
	a.Settings().SetTheme(theme.SqrTheme{})
	window.CreateMainWindow(a).ShowAndRun()
}