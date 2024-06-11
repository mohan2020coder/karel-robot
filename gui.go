package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
)

const cellSize = 50

func drawWorld(w *World, app fyne.App) fyne.CanvasObject {
    grid := canvas.NewGrid(w.cols, w.rows)
    for y := range grid.Rows {
        for x := range grid.Columns {
            cell := fyne.NewContainer(
                container.NewCenter(
                    fyne.NewCanvas(
                        func(paint canvas.Paint) {
                            paint.FillRect(fyne.NewColor("white"), fyne.NewSize(cellSize, cellSize))
                            if w.HasBeeper(Position{X: x, Y: y}) {
                                paint.Circle(fyne.NewColor("gold"), Position{X: cellSize / 2, Y: cellSize / 2}, cellSize/4)
                            }
                            // Draw robot based on its position (w.robot
                              // Draw robot based on its position (w.robot)
                              if x == w.robot.X && y == w.robot.Y {
                                  robotCenter := fyne.NewPos(float32(x*cellSize) + cellSize/2, float32(y*cellSize) + cellSize/2)
                                  paint.Circle(fyne.NewColor("blue"), robotCenter, cellSize/3)
                                  // You can also draw an arrow or other indicator to show robot's direction
                              }
                          },
                      ),
                  ),
              )
              grid.SetCell(x, y, cell)
          }
      }
      return grid
  }

  func createInterface(w *World, app fyne.App) *fyne.Container {
      input := fyne.NewEntry()
      input.SetPlaceHolder("Enter command (Move, PickBeeper, etc.)")
      runButton := fyne.NewButton("Run", func() {
          command := input.Text
          // Parse the command and call the corresponding World method (w.Move(), etc.)
          var err error
          var success bool
          switch command {
          case "Move":
              success, err = w.Move()
          case "PickBeeper":
              // Implement logic for PickBeeper()
          case "PutBeeper":
              // Implement logic for PutBeeper()
          case "TurnLeft":
              w.TurnRobot(TurnLeft)
          case "TurnRight":
              w.TurnRobot(TurnRight)
          default:
              // Handle invalid command
          }
          if err != nil {
              fyne.ShowNotification("Error: " + err.Error())
          } else if !success {
              fyne.ShowNotification("Action failed (e.g., hit a wall)")
          }
          // Refresh the GUI after each command execution
          canvas.Refresh(app.Driver().AllWindows()[0])
      })
      return container.NewVBox(drawWorld(w, app), container.NewHBox(input, runButton))
  }
