package component

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/usesqr/vault/db"
	"github.com/usesqr/vault/db/model"
)

func CreatePasswordList(win fyne.Window, entries *[]*model.Password) fyne.CanvasObject {
	res := db.DB.
		Order("created_at DESC").
		Find(&entries)

	if res.Error != nil {
		dialog.ShowError(res.Error, win)
	}

	label := widget.NewLabel("Select a password to view")
	hbox := container.NewHBox(container.NewCenter(label))

	list := widget.NewList(
		func() int {
			return len(*entries)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.AccountIcon()), widget.NewLabel("Template Object"))
		},
		nil,
	)
	list.UpdateItem = func(id widget.ListItemID, item fyne.CanvasObject) {
		name := (*entries)[id].Name
		username := (*entries)[id].Username

		if name == "" {
			name = "Unnamed "
		}
		if username == "" {
			username = "No username provided"
		}

		formattedName := fmt.Sprintf("%s\n%s", name, username)
		item.(*fyne.Container).Objects[1].(*widget.Label).SetText(formattedName)
		list.SetItemHeight(id, 50)
	}
	list.OnSelected = func(id widget.ListItemID) {
		hbox.Objects = []fyne.CanvasObject{CreatePasswordDetail(win, (*entries)[id])}
	}

	return container.NewHSplit(list, hbox)
}
