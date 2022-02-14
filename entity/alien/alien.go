package alien

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

//Choose next random location for an alien
func ChooseLocation(w WorldUseCase, a *Alien, s int64) (string, error) {
	return w.ProvideRandomCity(s)
}

//Move alien with the random direction
func (a *Alien) Move(w WorldUseCase, s int64, c chan string) (string, error) {
	return w.GetRandomNeighbor(a.Location, s)
}
