package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "github.com/mohan2020coder/karel-robot/world"
    "github.com/mohan2020coder/karel-robot/actions"
    "github.com/mohan2020coder/karel-robot/gui"
)

func main() {
    world := NewWorld(10, 10) // Create a world with 10 rows and 10 columns
    world.PutBeeper(Position{X: 5, Y: 3}) // Place a beeper at position (5, 3)

    a := app.New()
    w := a.NewWindow("Karel the Robot")
    w.SetContent(createInterface(world, a))
    w.Resize(fyne.NewSize(600, 600))
    w.ShowAndRun()
}
