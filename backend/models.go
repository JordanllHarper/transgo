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
	Coordinates
	Status
	TransportType
}

// Coordinates for locating on an x and y axis
type Coordinates struct{ x, y int }

// Information for a particular node
type NodeInformation struct {
	id   int
	name string
	Coordinates
	NodeType
}

// Singular destination in the transport map
// Stores neighbours as ids
type Node struct {
	neighbours []int
	NodeInformation
}

type JourneyStatus int

const (
	NotStarted JourneyStatus = iota
	InProgress
	Finished
)

// Singular instances of a Transport going between a source and destination Node.
type Journey struct {
	id                 int
	Source             Node
	Destination        Node
	percentageProgress float32
	JourneyStatus
	Transport
}
