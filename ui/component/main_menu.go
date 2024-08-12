package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/usesqr/vault/ui/dialog"
)

func CreateMainMenu(win fyne.Window) (menu *fyne.MainMenu) {
	newMenuItem := fyne.NewMenuItem("New password", func() { dialog.ShowPasswordCreateDialog(win) })
	newMenuItem.Icon = theme.ContentAddIcon()
	fileMenu := fyne.NewMenu("File", newMenuItem)

	aboutMenuItem := fyne.NewMenuItem("About", func() { dialog.ShowAboutDialog(win) })
	aboutMenuItem.Icon = theme.InfoIcon()
	helpMenu := fyne.NewMenu("Help", aboutMenuItem)

	return fyne.NewMainMenu(fileMenu, helpMenu)
}
