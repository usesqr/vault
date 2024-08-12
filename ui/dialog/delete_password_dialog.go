package dialog

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/usesqr/vault/db"
	"github.com/usesqr/vault/db/model"
	"github.com/usesqr/vault/global"
)

func ShowPasswordDeleteDialog(win fyne.Window, pass *model.Password) {
	name := pass.Name
	if name == "" {
		name = "Untitled"
	}
	confirm := dialog.NewConfirm(
		"Confirm deletion",
		fmt.Sprintf("Are you sure you'd like to permanently delete the password %q?", pass.Name),
		func(delete bool) {
			res := db.DB.Delete(pass)
			if res.Error != nil {
				dialog.ShowError(res.Error, win)
				return
			}

			dialog.ShowInformation("Success", "The password was successfully deleted.", win)

			db.DB.
				Order("created_at DESC").
				Find(&global.Passwords)
		},
		win,
	)
	confirm.Show()
}
