package component

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2"
	"github.com/usesqr/vault/ui/dialog"
)

func CreateMainMenu(win fyne.Window) (menu *fyne.MainMenu) {
	fileMenu := fyne.NewMenu("File")

	aboutMenuItem := fyne.NewMenuItem("About", func() { dialog.ShowAboutDialog(win) })
	aboutMenuItem.Icon = theme.InfoIcon()
	helpMenu := fyne.NewMenu("Help", aboutMenuItem)
	
	return fyne.NewMainMenu(fileMenu, helpMenu)
}