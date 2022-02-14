package world

import (
	"errors"
)

var WorldDirections = map[string]int{"north": 0, "south": 1, "east": 2, "west": 3}
var Sides = []string{"north", "south", "east", "west"}

//City is a node within graph
type City struct {
	Name       string
	Aliens     []string
	Directions []string
}

//World is technically a graph, allows for a random alien dislocation
type World struct {
	//List of all available cities
	Cities map[string]*City
}

//Create a constructor function for the Vertex
func NewCity(name string) City {
	return City{
		Name:       name,
		Directions: make([]string, 4),
	}
}

//Add neighbor d - direction, n - neighbor
func (c *City) AddNeighbor(d string, n City) error {
	wp := WorldDirections[d]

	if c.Directions[wp] != "" {
		return errors.New("there is a neighbor at this direction")
	}
	c.Directions[wp] = n.Name
	return nil
}

//Create a constructor for the graph
func NewWorld() World {
	return World{
		Cities: make(map[string]*City),
	}
}

//Populate new city to the world
//Get rid of the append, if performance are the priority, could use sync.pool for example
func (w *World) AddCity(c *City) {
	w.Cities[c.Name] = c
}

//Remove city after alien atack
func (w *World) RemoveCity(c *City) {
	delete(w.Cities, c.Name)
}
