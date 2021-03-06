package alien

import (
	"errors"
	"math/rand"

	"github.com/VladimirDemidov/alien-attack/entity/world"
)

//sync Mutex probably not the best solution, channels would be better solution
type Alien struct {
	Name     string
	Location *world.City
}

var WorldDirections = []string{"north", "south", "east", "west"}

//Create a new alien
func NewAlien(n string) *Alien {
	alien := &Alien{
		Name: n,
	}
	return alien
}

//Choose a random location for an alien
//This function is too long and some side effects are happening here,
//would be great to refactor it
func ChooseLocation(w *world.World, a *Alien, r int64) (*world.City, error) {
	rand.Seed(r)
	randomCityName := w.AvailableCities[rand.Intn(len(w.AvailableCities))]
	if city, ok := w.Cities[randomCityName.Name]; ok {
		//Append makes extra memory allocation, could be replaced with sync.pool
		city.Aliens = append(city.Aliens, a.Name)
		a.Location = city
		return city, nil
	}
	return nil, errors.New("City does not exists")
}

//Move alien with the random direction, from current city
func (a *Alien) Move(w *world.World, r int64, c chan string) error {
	rand.Seed(r)
	newDir := world.Sides[rand.Intn(len(world.Sides))]
	newDirIndex := world.WorldDirections[newDir]

	//This is way to huge lock, normally should not happen
	if nextCity, ok := a.Location.Directions[newDirIndex]; ok {
		if nc, ok := w.Cities[nextCity.Name]; ok {
			if len(nc.Aliens) < 1 {
				nc.Aliens = append(nc.Aliens, a.Name)
			} else {
				c <- a.Name
			}
			return nil
		}
		return nil
	}
	return errors.New("No such direction")
}
