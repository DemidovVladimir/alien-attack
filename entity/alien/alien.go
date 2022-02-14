package alien

import "fmt"

type Alien struct {
	Name     string
	Location string
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

//Choose a random location for an alien
//Adds string value to the w.InitialWorld
//Set string value to the a.Location
//Adds string value to the c.Aliens
func ChooseLocation(w WorldUseCase, a *Alien, s int64) error {
	city, err := w.ProvideRandomCity(s)
	if err != nil {
		return err
	}
	a.Location = city
	return nil
}

//Move alien with the random direction, from current city
//Use WorldUseCase in order to decouple things
func (a *Alien) Move(w WorldUseCase, s int64, c chan string) error {
	city, err := w.ProvideRandomCity(s)
	fmt.Println("Made a move to the", city)
	if err != nil {
		return err
	}
	return nil
}
