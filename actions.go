package main

type Direction int

const (
    North Direction = iota
    East
    South
    West
)

type TurnDirection int

const (
    TurnLeft TurnDirection = iota
    TurnRight
)

func (w *World) Move() (bool, error) {
    return w.MoveRobot(North)
}

// Implement functions for other actions: PickBeeper(), PutBeeper(), TurnLeft(), TurnRight()
// Use w.World methods to interact with the world
