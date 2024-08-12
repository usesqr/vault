package dialog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/usesqr/vault/db"
	"github.com/usesqr/vault/db/model"
	"github.com/usesqr/vault/global"
	"gorm.io/gorm"
)

func ShowPasswordUpdateDialog(win fyne.Window, pass *model.Password) {
	name := widget.NewEntry()
	name.SetPlaceHolder("App/site name")
	name.Text = pass.Name
	username := widget.NewEntry()
	username.SetPlaceHolder("Username")
	username.Text = pass.Username
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")
	password.Text = pass.Password
	items := []*widget.FormItem{
		widget.NewFormItem("", name),
		widget.NewFormItem("", username),
		widget.NewFormItem("", password),
	}
	dialog.ShowForm("Update a password", "Done", "Cancel", items, func(ok bool) {
		if !ok {
			return
		}

		res := db.DB.Save(&model.Password{
			Model: gorm.Model{
				ID: pass.ID,
			},
			Name:     name.Text,
			Username: username.Text,
			Password: password.Text,
		})

		if res.Error != nil {
			dialog.ShowError(res.Error, win)
		}

		passes := global.Passwords
		db.DB.
			Order("created_at DESC").
			Find(&passes)
	}, win)
}
