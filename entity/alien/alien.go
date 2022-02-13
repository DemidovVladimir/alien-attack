package alien

import (
	"math/rand"
	"time"

	"github.com/VladimirDemidov/alien-attack/entity/world"
)

type Alien struct {
	Name     string
	Status   string
	Location *world.City
}

var WorldDirections = []string{"north", "south", "east", "west"}

//All available aliens are here, move to the root
var LandedAliens = make(map[string]*Alien)

//Create a new alien
func NewAlien(n string) *Alien {
	alien := &Alien{
		Name: n,
	}
	LandedAliens[alien.Name] = alien
	return alien
}

//Choose a random location for an alien
//This function is too long and some side effects are happening here,
//would be great to refactor it
func ChooseLocation(w *world.World, a *Alien) *world.City {
	rand.Seed(time.Now().Unix())
	city := w.Cities[rand.Intn(len(w.Cities))]

	if len(city.Aliens) < 1 && w.Cities[city.PostCode-1] != nil {
		//Append makes extra memory allocation, could be replaced with sync.pool
		city.Aliens = append(city.Aliens, a)
		a.Location = city
		return city
	} else {
		a.Battle(city, w)
		return nil
	}
}

//Battle between aliens in the city
func (a *Alien) Battle(c *world.City, w *world.World) {
	// Kill all aliens that are currently in the same city
	for _, alien := range c.Aliens {
		if competitor, ok := alien.(*Alien); ok {
			competitor.Die()
		}
	}
	a.Die()
	w.RemoveCity(c)
}

//Move alien with the random direction, from current city
func (a *Alien) Move(w *world.World) {
	rand.Seed(time.Now().Unix())
	nextCity := a.Location.Directions[rand.Intn(len(a.Location.Directions))]
	if len(nextCity.Aliens) < 1 && nextCity != nil {
		a.Location.Aliens = nil
		a.Location = nextCity
		a.Location.Aliens = append(a.Location.Aliens, a)
	} else {
		a.Battle(nextCity, w)
	}
}

//Kill involved alien
func (a *Alien) Die() {
	delete(LandedAliens, a.Name)
}
