package world

import (
	"errors"
)

var WorldDirections = map[string]int{"north": 0, "south": 1, "east": 2, "west": 3}
var Sides = []string{"north", "south", "east", "west"}

//City is a node within graph
type City struct {
	Name   string
	Aliens []string
	//Think slice would fit better here
	Directions map[int]*City
}

//World is technically a graph, allows for a random alien dislocation
type World struct {
	//List of all available cities
	Cities          map[string]*City
	AvailableCities []*City
}

//Create a constructor function for the Vertex
func NewCity(name string) *City {
	return &City{
		Name:       name,
		Directions: map[int]*City{},
	}
}

//Add neighbor d - direction, n - neighbor, c - origin city
func (c *City) AddNeighbor(d int, n *City) error {
	if c.Directions[d] != nil {
		return errors.New("there is a neighbor at this direction")
	}
	c.Directions[d] = n
	return nil
}

//Create a constructor for the graph
func NewWorld() *World {
	return &World{
		Cities: make(map[string]*City),
	}
}

//Populate new city to the world
//Get rid of the append, if performance are the priority, could use sync.pool for example
func (w *World) AddCity(c *City) {
	w.Cities[c.Name] = c
	w.AvailableCities = append(w.AvailableCities, c)
}

//Remove city after alien atack
func (w *World) RemoveCity(c *City) {
	delete(w.Cities, c.Name)
}
