package utils

import (
	"time"

	"fyne.io/fyne/v2"
)

func ShowNotification(a fyne.App) {
	time.Sleep(time.Second * 2)
	a.SendNotification(fyne.NewNotification("Water Remember", "Is time to drink!!!"))
}
