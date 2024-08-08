package dialog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
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
	dialog.ShowForm("Create a new password entry", "Done", "Cancel", items, func(b bool) {}, win)
}