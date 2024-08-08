package dialog

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/usesqr/vault/consts"
)

func ShowAboutDialog(win fyne.Window) {
	dialog.ShowInformation(
		consts.FullName,
		fmt.Sprintf("Version %s\n%s", consts.Version, consts.GitHub),
		win,
	)
}
