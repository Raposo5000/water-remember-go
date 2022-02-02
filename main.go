package main

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/Raposo5000/waterremember/utils"
)

type Notification struct {
	Title, Content string
}

func main() {
	application := app.New()
	a := app.NewWithID("water.remember")
	window := application.NewWindow("Water Remember")

	// CONFIGURATION
	seconds_drink := 1200 // 20 minutes
	ml_per_drink := 250   // 250 ml

	// TIME TO NEXT DRINK
	timer_to_drink_in_seconds := seconds_drink
	text_label_time_to_drink := "Time for next drink: " + strconv.Itoa(timer_to_drink_in_seconds)
	label_time := widget.NewLabel(text_label_time_to_drink)

	// WATER CONSUMED
	drinked := 0
	text_label_drinked := "Drink quantity: " + strconv.Itoa(drinked) + "ml"
	label_qnt_drinked := widget.NewLabel(text_label_drinked)

	is_time_to_drink := false

	var ticker *time.Ticker

	// FUNCTIONS
	stopTimer := func() {
		ticker.Stop()
	}

	startTimer := func() {
		ticker = time.NewTicker(time.Second * 1)
		if is_time_to_drink == false {
			go func() {
				for range ticker.C {
					timer_to_drink_in_seconds--
					label_time.SetText("Time for next drink: " + strconv.Itoa(timer_to_drink_in_seconds))
					fmt.Println("Seconds: ", timer_to_drink_in_seconds)
					if timer_to_drink_in_seconds == 0 {
						label_time.SetText("IT'S TIME TO DRINK!!!!")
						go utils.ShowNotification(a)
						window.RequestFocus()
						stopTimer()
						is_time_to_drink = true
					}
				}
			}()
		}
	}

	drink := func() {
		drinked = drinked + ml_per_drink
		label_qnt_drinked.SetText("Drink quantity: " + strconv.Itoa(drinked) + "ml")
		timer_to_drink_in_seconds = seconds_drink
		// ticker = time.NewTicker(time.Second * 1)
		is_time_to_drink = false
		startTimer()
	}

	// SET WIDGETS
	window.SetContent(container.NewVBox(
		label_time,
		label_qnt_drinked,
		widget.NewButton("START!", func() {
			fmt.Println("START!")
			// timer = true
			defer startTimer()
		}),
		widget.NewButton("DRINK", func() {
			drink()
		}),
		widget.NewButton("STOP", func() {
			fmt.Println("STOP =/")
			stopTimer()
		}),
	))

	window.ShowAndRun()
}
