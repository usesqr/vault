package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/usesqr/vault/ui/credits"
	"github.com/usesqr/vault/ui/dialog"
)

func CreateMainMenu(win fyne.Window) (menu *fyne.MainMenu) {
	newMenuItem := fyne.NewMenuItem("New password", func() { dialog.ShowPasswordCreateDialog(win) })
	newMenuItem.Icon = theme.ContentAddIcon()
	fileMenu := fyne.NewMenu("File", newMenuItem)

	licensesMenuItem := fyne.NewMenuItem("Licenses", func() {
		credits.CreditsWindow(fyne.CurrentApp(), fyne.NewSize(800, 400)).Show()
	})
	licensesMenuItem.Icon = theme.FileTextIcon()
	aboutMenuItem := fyne.NewMenuItem("About", func() { dialog.ShowAboutDialog(win) })
	aboutMenuItem.Icon = theme.InfoIcon()
	helpMenu := fyne.NewMenu("Help", licensesMenuItem, aboutMenuItem)

	return fyne.NewMainMenu(fileMenu, helpMenu)
}
