package window

import (
	"os"
	"strings"

	"fyne.io/fyne/v2"
	fyneDialog "fyne.io/fyne/v2/dialog"
	"github.com/usesqr/vault/consts"
	"github.com/usesqr/vault/crypto"
	"github.com/usesqr/vault/db"
	"github.com/usesqr/vault/global"
	"github.com/usesqr/vault/ui/component"
	"github.com/usesqr/vault/ui/dialog"
)

func CreateMainWindow(a fyne.App) fyne.Window {
	w := a.NewWindow(consts.FullName)
	w.Resize(fyne.Size{Width: 800, Height: 600})
	w.SetMainMenu(component.CreateMainMenu(w))

	dbConnected := make(chan bool)

	if crypto.DoesSaltFileExist() {
		dialog.ShowUnlockDialog(w, func(pass string, ok bool) {
			if !ok {
				os.Exit(1)
			}

			err := db.Connect(pass, false)
			if err != nil {
				if strings.Contains(err.Error(), "file is not a database") {
					dialog := fyneDialog.NewInformation(
						"Could not decrypt database",
						"Check that you typed the right master password.",
						w,
					)
					dialog.SetDismissText("Quit app")
					dialog.SetOnClosed(func() { os.Exit(1) })
					dialog.Show()
				} else {
					panic(err)
				}
				return
			}
			dbConnected <- true
		})
	} else {
		dialog.ShowFirstStartDialog(w, func(pass string, ok bool) {
			if !ok {
				os.Exit(1)
			}

			err := db.Connect(pass, false)
			if err != nil {
				panic(err)
			}
			dbConnected <- true
		})
	}

	go func() {
		<-dbConnected
		toolbar := component.CreateToolbar(w).(*fyne.Container)
		toolbar.Add(component.CreatePasswordList(w, global.Passwords))
		w.SetContent(toolbar)
	}()

	return w
}
