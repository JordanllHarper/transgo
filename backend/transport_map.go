package main

import (
	"errors"
	"strconv"
	"strings"
)

// need a way to describe nodes and their connections
// this can be a file format of some kind
// thinking describing this as so:

/*
id name coordinates node_type
another_id name coordinates node_type
*new line* to distinguish nodes from connections
id another_id <- represents connection
*/
type TransportGraph Node

func parse_node_connections(connections_section string, nodes map[int]NodeInformation) (TransportGraph, error) {
	// TODO: Implement
	return TransportGraph{}, nil
}

func parse_node_type(node_type_str string) (NodeType, error) {
	node_type := NodeType(-1)
	switch node_type_str {
	case "train_station":
		node_type = TrainStation
	case "airport":
		node_type = TrainStation
	default:
		return node_type, errors.New("Invalid node type")
	}
	return node_type, nil
}
func parse_coords(coordinate_x string, coordinate_y string) (Coordinates, error) {

	x, err := strconv.Atoi(coordinate_x)
	if err != nil {
		return Coordinates{}, err
	}

	y, err := strconv.Atoi(coordinate_y)
	if err != nil {
		return Coordinates{}, err
	}

	return Coordinates{x: x, y: y}, nil

}
func parse_node_information(node_section string) (map[int]NodeInformation, error) {
	node_information := map[int]NodeInformation{}
	lines := strings.Split(node_section, "\n")
	for _, line := range lines {

		sections := strings.Split(line, " ")
		id_str := sections[0]

		id, err := strconv.Atoi(id_str)
		if err != nil {
			return map[int]NodeInformation{}, err
		}

		name := sections[1]
		coordinates_x := sections[2]
		coordinates_y := sections[3]
		coords, err := parse_coords(coordinates_x, coordinates_y)

		if err != nil {
			return map[int]NodeInformation{}, err
		}

		node_type, err := parse_node_type(sections[4])
		if err != nil {
			return map[int]NodeInformation{}, err
		}

		node_information[id] = NodeInformation{id: id, name: name, Coordinates: coords, NodeType: node_type}
	}

	return node_information, nil
}
func parse_transport_graph(input string) (TransportGraph, error) {
	split := strings.Split(input, "\n\n")
	if len(split) < 2 {
		return TransportGraph{}, errors.New("Invalid file input")
	}

	node_info, err := parse_node_information(split[0])

	if err != nil {
		return TransportGraph{}, err
	}

	node_connections, err := parse_node_connections(split[1], node_info)

	if err != nil {
		return TransportGraph{}, err
	}

	return node_connections, nil
}
