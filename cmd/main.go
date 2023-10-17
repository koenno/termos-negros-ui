package main

import (
	"github.com/koenno/termos-negros-ui/app"
	"github.com/koenno/termos-negros-ui/data"
)

func main() {
	theApp := app.New()

	menu := data.NewMenu()
	d, err := menu.GetData()
	if err != nil {
		theApp.ShowError(err)
	} else {
		theApp.SetData(d)
		theApp.ShowData()
	}

	theApp.Run()
}
