package dialog

import (
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ShowUnlockDialog(win fyne.Window, callback func(string, bool)) {
	password := widget.NewPasswordEntry()
	password.Validator = validation.NewRegexp(".+", "Field is required")
	items := []*widget.FormItem{
		widget.NewFormItem("", password),	
	}
	dialog.ShowForm("Enter your master password.", "Done", "Cancel", items, func(b bool) {
		callback(password.Text, b)
	}, win)
}