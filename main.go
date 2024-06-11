package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Karel struct {
	X, Y int
	Dir  Direction
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

var dirSymbols = [...]rune{'^', '>', 'v', '<'}

func (k *Karel) TurnLeft() {
	k.Dir = (k.Dir + 3) % 4
}

func (k *Karel) TurnRight() {
	k.Dir = (k.Dir + 1) % 4
}

func (k *Karel) Move() {
	switch k.Dir {
	case North:
		k.Y--
	case South:
		k.Y++
	case West:
		k.X--
	case East:
		k.X++
	}
}

func main() {
	myApp := app.New()
	win := myApp.NewWindow("Karel the Robot")

	karel := &Karel{X: 0, Y: 0, Dir: East}

	karelCanvas := canvas.NewText(string(dirSymbols[karel.Dir]), color.Black)

	grid := container.NewGridWithColumns(2,
		canvas.NewRectangle(color.White),
		container.NewVBox(
			karelCanvas,
			widget.NewButton("Move", func() {
				karel.Move()
				karelCanvas.Text = string(dirSymbols[karel.Dir])
				win.Canvas().Refresh(karelCanvas)
			}),
			widget.NewButton("Turn Left", func() {
				karel.TurnLeft()
				karelCanvas.Text = string(dirSymbols[karel.Dir])
				win.Canvas().Refresh(karelCanvas)
			}),
			widget.NewButton("Turn Right", func() {
				karel.TurnRight()
				karelCanvas.Text = string(dirSymbols[karel.Dir])
				win.Canvas().Refresh(karelCanvas)
			}),
		),
	)

	win.SetContent(grid)
	win.Resize(fyne.NewSize(400, 400))
	win.ShowAndRun()
}
