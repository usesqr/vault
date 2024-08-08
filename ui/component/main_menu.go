package component

import (
	"fyne.io/fyne/v2"
)

func CreateMainMenu() (menu *fyne.MainMenu) {
	fileMenu := fyne.NewMenu("File")
	return fyne.NewMainMenu(fileMenu)
}