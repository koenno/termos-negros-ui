package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	fyneapp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/koenno/termos-negros/domain"
)

type App struct {
	gui        fyne.App
	mainWindow fyne.Window
	data       domain.Menu
}

func New() *App {
	gui := fyneapp.New()
	mainWindow := gui.NewWindow("Menu")
	return &App{
		gui:        gui,
		mainWindow: mainWindow,
	}
}

func (a *App) SetData(data domain.Menu) {
	a.data = data
}

func (a *App) ShowMenu(menu domain.Menu) {
	// progress := widget.NewProgressBarInfinite()
	// // progress.Resize(fyne.NewSize(20.0, 5.0))
	// content := container.New(layout.NewCenterLayout(), progress)
	// a.mainWindow.SetContent(content)

	filterButton := widget.NewButton("hide outdated", func() { fmt.Println("tapped text button") })
	buttonsContainer := container.NewVBox(filterButton)

	dataView := menuView(menu)

	mainBox := container.NewBorder(nil, buttonsContainer, nil, nil, dataView)

	a.mainWindow.SetContent(mainBox)
	a.mainWindow.SetFullScreen(false)
	a.mainWindow.Resize(fyne.NewSize(400.0, 600.0))
}

func (a *App) ShowError(err error) {
	content := widget.NewTextGridFromString(err.Error())
	a.mainWindow.SetContent(content)
}

func (a *App) Run() {
	a.mainWindow.Show()
	a.gui.Run()
}
