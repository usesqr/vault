package dialog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/usesqr/vault/db"
	"github.com/usesqr/vault/db/model"
	"github.com/usesqr/vault/global"
)

func ShowPasswordCreateDialog(win fyne.Window) {
	name := widget.NewEntry()
	name.SetPlaceHolder("App/site name")
	username := widget.NewEntry()
	username.SetPlaceHolder("Username")
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")
	items := []*widget.FormItem{
		widget.NewFormItem("", name),	
		widget.NewFormItem("", username),	
		widget.NewFormItem("", password),	
	}
	dialog.ShowForm("Create a new password entry", "Done", "Cancel", items, func(ok bool) {
		if !ok { return }

		pass := &model.Password{
			Name: name.Text,
			Username: username.Text,
			Password: password.Text,
		}
		res := db.DB.Create(pass)

		if res.Error != nil {
			dialog.ShowError(res.Error, win)
		}

		passes := global.Passwords
		*passes = append(*passes, pass)
	}, win)
}