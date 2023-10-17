package main

import (
	"github.com/koenno/termos-negros-ui/app"
	"github.com/koenno/termos-negros-ui/data"
)

func main() {
	// myApp := app.New()
	// myApp.
	// 	myWindow := myApp.NewWindow("Menu")
	// black := color.NRGBA{R: 0, G: 0, B: 0, A: 255}

	// text1 := canvas.NewText(`Zupka pomidorowa z pomidorów pelati na wywarze warzywnym z marchewki, pietruszki, selera (9), pora, koperku, czosnku z makaronem (1,3) (dla wersji bezglutenowej makaron z mąki ryżowej, kukurydzianej i quinoa) Racuchy z bananami z mąką orkiszową pełnoziarnistą, z cynamonem i kardamonem :) (1,3) micha soczystych jabłek do pochrupania :)`, black)
	// // text2 := canvas.NewText("There", black)
	// // text2.Move(fyne.NewPos(2000, 500))
	// card0 := widget.NewCard("śniadanie", "", text1)
	// card1 := widget.NewCard("obiad", "", text1)
	// card2 := widget.NewCard("podwieczorek", "", text1)
	// day1Content := container.New(layout.NewGridLayout(1), card0, card1, card2)
	// card := widget.NewCard("poniedziałek", "", day1Content)
	// content := container.New(layout.NewGridLayout(2), card)

	// myWindow.SetContent(content)
	// myWindow.ShowAndRun()

	menu := data.NewMenu()
	dataPipe, errPipe := menu.GetData()
	theApp := app.New()

	go func() {
		for err := range errPipe {
			theApp.ShowError(err)
		}
	}()

	theApp.ShowMenu(dataPipe)
	theApp.Run()
}
