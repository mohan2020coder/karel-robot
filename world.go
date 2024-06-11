package main

import (
    "fyne.io/fyne/v2/canvas"
)

type World struct {
    rows, cols int
    beepers    map[Position]bool
    robot      Position
}

type Position struct {
    X, Y int
}

func NewWorld(rows, cols int) *World {
    return &World{
        rows:   rows,
        cols:   cols,
        beepers: make(map[Position]bool),
        robot:   Position{0, 0}, // Initial position
    }
}

func (w *World) HasBeeper(p Position) bool {
    return w.beepers[p]
}

func (w *World) PutBeeper(p Position) {
    w.beepers[p] = true
}

func (w *World) PickBeeper(p Position) bool {
    if !w.HasBeeper(p) {
        return false
    }
    delete(w.beepers, p)
    return true
}

func (w *World) MoveRobot(dir Direction) (bool, error) {
    newPos := w.robot.Add(dir)
    if newPos.X < 0 || newPos.X >= w.cols || newPos.Y < 0 || newPos.Y >= w.rows {
        return false, nil // Hit a wall
    }
    w.robot = newPos
    return true, nil
}

func (w *World) TurnRobot(dir TurnDirection) {
    // Implement logic to update robot direction based on TurnDirection
}
