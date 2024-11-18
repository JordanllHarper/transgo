package main

type Status int

const (
	Travelling Status = iota
	Transferring
	Unused
	Emergency
)

type TransportType int

const (
	Train TransportType = iota
	Plane
)

type NodeType int

const (
	TrainStation = iota
	Airport
)

// Singular transport entity that can go from a source to a destination
type Transport struct {
	id             int
	name           string
	travelTimeInMs int
	Status
	TransportType
}

// Coordinates for locating on an x and y axis
type Coordinates struct {
	x int
	y int
}

// Singular destination in the transport map
type Node struct {
	id         int
	name       string
	neighbours []Node
	Coordinates
	NodeType
}

type JourneyStatus int

const (
	NotStarted JourneyStatus = iota
	InProgress
	Finished
)

type Journey struct {
	id                 int
	Source             Node
	Destination        Node
	percentageProgress float32
	JourneyStatus
	Transport
}
