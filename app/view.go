package app

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/koenno/termos-negros/domain"
)

func mealView(meal domain.Meal) fyne.CanvasObject {
	name := widget.NewLabel(meal.Name)
	name.Alignment = fyne.TextAlignCenter
	name.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	ingredients := widget.NewLabel(meal.Ingredients)
	ingredients.Wrapping = fyne.TextWrapWord
	view := container.New(layout.NewVBoxLayout(), name, ingredients)
	return view
}

func dayView(dayMenu domain.DayMenu) fyne.CanvasObject {
	day := canvas.NewText(dayMenu.Date.Format(time.DateOnly), color.Black)
	day.Alignment = fyne.TextAlignTrailing
	day.TextStyle = fyne.TextStyle{Bold: true}

	var mealsView []fyne.CanvasObject
	for _, m := range dayMenu.Meals {
		mealsView = append(mealsView, mealView(m))
	}

	c := container.New(layout.NewVBoxLayout(), mealsView...)
	view := widget.NewCard(dayMenu.Date.Format("Monday"), dayMenu.Date.Format("02.01.2006"), c)
	return view
}

func menuView(menu domain.Menu) fyne.CanvasObject {
	box := container.NewVBox()
	for _, dayMenu := range menu {
		box.Add(dayView(dayMenu))
	}
	scrollContent := container.NewScroll(box)
	return scrollContent
}
