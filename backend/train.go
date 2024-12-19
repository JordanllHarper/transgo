package main

import "gorm.io/gorm"

type Status string

const (
	Travelling   Status = "Travelling"
	Transferring        = "Transferring"
	Unused              = "Unused"
	Emergency           = "Emergency"
)


type TrainEntity struct {
	gorm.Model
	Train
}
type Train struct {
	Name           string
	TravelTimeInMs int
	Coordinates
	Status
}
