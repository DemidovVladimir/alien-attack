package alien

import "sync"

type Alien struct {
	Name     string
	Location string
}

type Swarm struct {
	Aliens sync.Map
	//World before invasion
	LandedAliens []string
}

//Code duplication for low coupling
var WorldDirections = map[string]int{"north": 0, "south": 1, "east": 2, "west": 3}
var Sides = []string{"north", "south", "east", "west"}

//Create a new alien
func NewAlien(n string) *Alien {
	return &Alien{
		Name: n,
	}
}

//Create a new swarm
func NewSwarm() *Swarm {
	return &Swarm{}
}

//Choose next random location for an alien
func ChooseLocation(w WorldUseCase, a *Alien, s int) (string, error) {
	return w.ProvideRandomCity(int64(s))
}

//Move alien with the random direction
func (a *Alien) Move(w WorldUseCase, s int) (string, error) {
	return w.GetRandomNeighbor(a.Location, int64(s))
}
