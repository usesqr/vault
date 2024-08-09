package window

import (
	"os"

	"fyne.io/fyne/v2"
	"github.com/usesqr/vault/consts"
	"github.com/usesqr/vault/crypto"
	"github.com/usesqr/vault/db"
	"github.com/usesqr/vault/ui/component"
	"github.com/usesqr/vault/ui/dialog"
)

func CreateMainWindow(a fyne.App) (w fyne.Window) {
	w = a.NewWindow(consts.FullName)
	w.Resize(fyne.Size{Width: 800, Height: 600})
	w.SetMainMenu(component.CreateMainMenu(w))

	w.SetContent(component.CreateToolbar(w))

	if crypto.DoesSaltFileExist() {
		dialog.ShowUnlockDialog(w, func(pass string, ok bool) {
			if !ok {
				os.Exit(1)
			}
	
			db.Connect(pass, false)
		})
	} else {
		dialog.ShowFirstStartDialog(w, func(pass string, ok bool) {
			if !ok {
				os.Exit(1)
			}
	
			db.Connect(pass, true)
		})
	}
	
	return
}