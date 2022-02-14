package world

import (
	"errors"
	"math/rand"
)

var WorldDirections = map[string]int{"north": 0, "south": 1, "east": 2, "west": 3}
var Sides = []string{"north", "south", "east", "west"}

//City
type City struct {
	Name       string
	Aliens     []string
	Directions []string
}

//Moving out from complexity
type World struct {
	//List of all available cities
	//Takes part in validation, easy to remove with bigO notation
	Cities map[string]*City
	//World before invasion
	InitialWorld []string
}

//Create a constructor function for City
func NewCity(name string) City {
	return City{
		Name: name,
		//There are only 4 parts currently available, maybe some other word have more
		//Not super flexible, as it uses Array and not Slice
		Directions: make([]string, 4),
	}
}

//Add neighbor d - direction, n - neighbor
//Get numeric value of the provided direction(north, south, etc)
//Set this numeric value as a key to the c.Directions and Name of the city(string) as a value
func (c *City) AddNeighbor(d string, n City) error {
	wp := WorldDirections[d]

	if c.Directions[wp] != "" {
		return errors.New("there is a neighbor at this direction")
	}
	c.Directions[wp] = n.Name
	return nil
}

//Create a constructor for the World
//Populates Name but no need for directions at this stage
func NewWorld() World {
	return World{
		Cities: make(map[string]*City),
	}
}

//Populate new city to the world
//Adds city to the map
//Adds city to the inocent world list
//Get rid of the append, if performance are the priority, could use sync.pool for example
func (w *World) AddCity(c *City) {
	w.Cities[c.Name] = c
	w.InitialWorld = append(w.InitialWorld, c.Name)
}

//Remove city after alien atack
//Remove from map, used further to validate left cities
func (w *World) RemoveCity(c *City) {
	delete(w.Cities, c.Name)
}

//Provide random city for alien landing and move
//Randomly picks string value from the w.InitialWorld list
//Picks this city name from the w.Cities map and validates if it still exists
func (w *World) ProvideRandomCity(s int64) (string, error) {
	if len(w.InitialWorld) < 1 {
		return "", errors.New("the world is empty")
	}
	rand.Seed(s)
	randomCityName := w.InitialWorld[rand.Intn(len(w.InitialWorld))]
	if c, ok := w.Cities[randomCityName]; ok {
		return c.Name, nil
	}
	return "", errors.New("city does not exists")
}
