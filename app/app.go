package app

import (
	"time"

	"fyne.io/fyne/v2"
	fyneapp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/koenno/termos-negros/domain"
	"github.com/koenno/termos-negros/filter"
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

func (a *App) ShowData() {
	// progress := widget.NewProgressBarInfinite()
	// // progress.Resize(fyne.NewSize(20.0, 5.0))
	// content := container.New(layout.NewCenterLayout(), progress)
	// a.mainWindow.SetContent(content)

	dataView := menuView(a.data)

	buttonsContainer := a.hideOutdatedButton()

	mainBox := container.NewBorder(nil, buttonsContainer, nil, nil, dataView)

	a.mainWindow.SetContent(mainBox)
	a.mainWindow.SetFullScreen(false)
	a.mainWindow.Resize(fyne.NewSize(400.0, 600.0))
}

func (a *App) hideOutdated() {
	now := time.Now()
	filteredMenu := filter.Since(now, a.data)
	dataView := menuView(filteredMenu)
	buttonsContainer := a.showAllButton()
	mainBox := container.NewBorder(nil, buttonsContainer, nil, nil, dataView)

	a.mainWindow.SetContent(mainBox)
}

func (a *App) showAll() {
	dataView := menuView(a.data)
	buttonsContainer := a.hideOutdatedButton()
	mainBox := container.NewBorder(nil, buttonsContainer, nil, nil, dataView)

	a.mainWindow.SetContent(mainBox)
}

func (a *App) ShowError(err error) {
	content := widget.NewTextGridFromString(err.Error())
	a.mainWindow.SetContent(content)
}

func (a *App) Run() {
	a.mainWindow.Show()
	a.gui.Run()
}

func (a *App) hideOutdatedButton() fyne.CanvasObject {
	filterButton := widget.NewButton("hide outdated", nil)
	filterButton.OnTapped = a.hideOutdated
	buttonsContainer := container.NewVBox(filterButton)
	return buttonsContainer
}

func (a *App) showAllButton() fyne.CanvasObject {
	filterButton := widget.NewButton("show all", nil)
	filterButton.OnTapped = a.showAll
	buttonsContainer := container.NewVBox(filterButton)
	return buttonsContainer
}
