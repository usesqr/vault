package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/usesqr/vault/db/model"
)

func createBoldLabel(text string) (label *widget.Label) {
	label = widget.NewLabel(text)
	label.TextStyle.Bold = true
	return
}

func CreatePasswordDetail(win fyne.Window, pass *model.Password) fyne.CanvasObject {
	return container.NewVBox(
		container.NewHBox(
			createBoldLabel("Name"),
			widget.NewLabel(pass.Name),
		),
		container.NewHBox(
			createBoldLabel("Username"),
			widget.NewLabel(pass.Username),
		),
		container.NewHBox(
			createBoldLabel("Password"),
			widget.NewLabel(pass.Password),
		),
	)
}