package world

import (
	"errors"
)

var WorldDirections = map[string]int{"north": 0, "south": 1, "east": 2, "west": 3}

//City is a node within graph
type City struct {
	Name   string
	Aliens []interface{}
	//Think slice would fit better here
	Directions map[int]*City
	PostCode   int
}

//World is technically a graph, allows for a random alien dislocation
type World struct {
	//List of all available cities
	Cities []*City
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

//Validate weather the neighbor stil exists
func (c *City) NeighborStillExists(d int, w *World) bool {
	return w.Cities[c.Directions[d].PostCode-1] != nil
}

//Create a constructor for the graph
func NewWorld() *World {
	return &World{}
}

//Populate new city to the world
//Get rid of the append, if performance are the priority, could use sync.pool for example
func (w *World) AddCity(c *City) {
	c.PostCode = len(w.Cities) + 1
	w.Cities = append(w.Cities, c)
}

//Remove city after alien atack
func (w *World) RemoveCity(c *City) {
	w.Cities[c.PostCode-1] = nil
}
