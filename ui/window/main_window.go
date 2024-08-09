package window

import (
	"os"

	"fyne.io/fyne/v2"
	"github.com/usesqr/vault/consts"
	"github.com/usesqr/vault/ui/component"
	"github.com/usesqr/vault/ui/dialog"
)

func CreateMainWindow(a fyne.App) (w fyne.Window) {
	w = a.NewWindow(consts.FullName)
	w.Resize(fyne.Size{Width: 800, Height: 600})
	w.SetMainMenu(component.CreateMainMenu(w))

	w.SetContent(component.CreateToolbar(w))
	dialog.ShowFirstStartDialog(w, func(pass string, ok bool) {
		if !ok {
			os.Exit(1)
		}

		println(pass)
	})

	return
}