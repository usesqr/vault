package component

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/usesqr/vault/db"
	"github.com/usesqr/vault/db/model"
)

func CreatePasswordList(win fyne.Window) fyne.CanvasObject {
	entries := new([]*model.Password)
	res := db.DB.
		Order("name").
		Find(&entries)
	
	if res.Error != nil {
		dialog.ShowError(res.Error, win)
	}

	label := widget.NewLabel("Select a password to view")
	hbox := container.NewHBox(label)

	list := widget.NewList(
		func() int {
			return len(*entries)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel("sadf"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			formattedName := fmt.Sprintf("%s: %s", (*entries)[id].Name, (*entries)[id].Username)
			item.(*fyne.Container).Objects[0].(*widget.Label).SetText(formattedName)
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		label.SetText((*entries)[id].Name)
	}
	list.OnUnselected = func(id widget.ListItemID) {
		label.SetText("Select a password to view")
	}

	return container.NewHSplit(list, container.NewCenter(hbox))
}