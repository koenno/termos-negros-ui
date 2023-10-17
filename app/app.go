package app

import (
	"fyne.io/fyne/v2"
	fyneapp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/koenno/termos-negros/domain"
)

type App struct {
	gui        fyne.App
	mainWindow fyne.Window
}

func New() App {
	gui := fyneapp.New()
	mainWindow := gui.NewWindow("Menu")
	return App{
		gui:        gui,
		mainWindow: mainWindow,
	}
}

func (a App) ShowMenu(dataPipe <-chan domain.DayMenu) {
	// progress := widget.NewProgressBarInfinite()
	// // progress.Resize(fyne.NewSize(20.0, 5.0))
	// content := container.New(layout.NewCenterLayout(), progress)
	// a.mainWindow.SetContent(content)

	box := container.NewVBox()
	menu := container.NewAdaptiveGrid(1, box)
	menu.Resize(fyne.NewSize(400.0, 600.0))
	for data := range dataPipe {
		box.Add(dayView(data))
	}

	scrollContent := container.NewScroll(menu)

	a.mainWindow.SetContent(scrollContent)
	a.mainWindow.SetFullScreen(false)
	a.mainWindow.Resize(fyne.NewSize(400.0, 600.0))
}

func (a App) ShowError(err error) {
	content := widget.NewTextGridFromString(err.Error())
	a.mainWindow.SetContent(content)
}

func (a App) Run() {
	a.mainWindow.Show()
	a.gui.Run()
}
