package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/usesqr/vault/ui/dialog"
)

func CreateToolbar(win fyne.Window) fyne.CanvasObject {
	t := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() { dialog.ShowPasswordCreateDialog(win) }),
	)

	return container.NewBorder(t, nil, nil, nil)
}