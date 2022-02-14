package world

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCity(t *testing.T) {
	test := "London"
	city := NewCity(test)
	assert.Equal(t, city.Name, test)
}

func TestAddCity(t *testing.T) {
	test := "London"
	city := NewCity(test)
	world := NewWorld()
	world.AddCity(&city)
	assert.Equal(t, len(world.Cities), 1)
}

func TestRemoveCity(t *testing.T) {
	test := "Tokio"
	neighbor := "Minato"
	city := NewCity(test)
	neighborCity := NewCity(neighbor)
	world := NewWorld()
	world.AddCity(&city)
	world.AddCity(&neighborCity)
	world.RemoveCity(&city)
	assert.Equal(t, neighborCity.Name, world.InitialWorld[1])
}

func TestNewWorld(t *testing.T) {
	world := NewWorld()
	assert.Equal(t, len(world.Cities), 0)
}

func TestAddNeighbor(t *testing.T) {
	test := "Tokio"
	neighbor := "Minato"
	city := NewCity(test)
	neighborCity := NewCity(neighbor)
	city.AddNeighbor("north", neighborCity)
	assert.Equal(t, city.Directions[0], neighbor)
}

func TestNeighborStillExists(t *testing.T) {
	test := "Tokio"
	neighbor := "Minato"
	city := NewCity(test)
	neighborCity := NewCity(neighbor)
	world := NewWorld()
	world.AddCity(&city)
	world.AddCity(&neighborCity)
	city.AddNeighbor("north", neighborCity)
	//Validate if we have a neighbor
	world.RemoveCity(&neighborCity)
	//Validate if our neighbor was defeated
	assert.Equal(t, 2, len(world.InitialWorld))
	assert.Nil(t, world.Cities[neighborCity.Name])
}

func TestAddNeighborFailure(t *testing.T) {
	test := "Tokio"
	neighbor1 := "Minato"
	neighbor2 := "Shinagava"
	city := NewCity(test)
	neighborCity1 := NewCity(neighbor1)
	neighborCity2 := NewCity(neighbor2)
	city.AddNeighbor("south", neighborCity1)
	err := city.AddNeighbor("south", neighborCity2)
	assert.Equal(t, err, errors.New("there is a neighbor at this direction"))
}

func TestProvideCity(t *testing.T) {
	var s int64
	test := "Tokio"
	neighbor := "Minato"
	city := NewCity(test)
	neighborCity := NewCity(neighbor)
	world := NewWorld()
	world.AddCity(&city)
	world.AddCity(&neighborCity)
	name, _ := world.ProvideRandomCity(s)
	assert.NotNil(t, name)
}

func TestProvideCityEmptyWorld(t *testing.T) {
	var s int64
	world := NewWorld()
	_, err := world.ProvideRandomCity(s)
	assert.Error(t, err)
}

func TestProvideCityWithRemovedCity(t *testing.T) {
	var s int64
	test := "Tokio"
	city := NewCity(test)
	world := NewWorld()
	world.AddCity(&city)
	world.RemoveCity(&city)
	_, err := world.ProvideRandomCity(s)
	assert.Error(t, err)
}

func TestAddAlienOrFight(t *testing.T) {
	test := "Tokio"
	city := NewCity(test)
	world := NewWorld()
	world.AddCity(&city)
	city.AddAlienOrFight("Gorm")
	assert.Equal(t, 1, len(city.Aliens))
}

func TestAddAlienOrFightError(t *testing.T) {
	test := "Tokio"
	city := NewCity(test)
	world := NewWorld()
	world.AddCity(&city)
	city.AddAlienOrFight("Gorm")
	err := city.AddAlienOrFight("Uber")
	assert.Error(t, err)
}

func TestGetCityByName(t *testing.T) {
	test := "Tokio"
	city := NewCity(test)
	world := NewWorld()
	world.AddCity(&city)
	c, _ := world.GetCityByName(test)
	assert.NotNil(t, c)
}

func TestGetCityByNameFail(t *testing.T) {
	test := "Tokio"
	world := NewWorld()
	_, err := world.GetCityByName(test)
	assert.Error(t, err)
}

func TestGetRandomNeighbor(t *testing.T) {
	var s int64
	test := "Tokio"
	neighbor1 := "Minato"
	neighbor2 := "London"
	neighbor3 := "Krakow"
	neighbor4 := "Capetown"
	city := NewCity(test)
	neighborCity1 := NewCity(neighbor1)
	neighborCity2 := NewCity(neighbor2)
	neighborCity3 := NewCity(neighbor3)
	neighborCity4 := NewCity(neighbor4)
	world := NewWorld()
	world.AddCity(&city)
	world.AddCity(&neighborCity1)
	world.AddCity(&neighborCity2)
	world.AddCity(&neighborCity3)
	world.AddCity(&neighborCity4)
	city.AddNeighbor("north", neighborCity1)
	city.AddNeighbor("south", neighborCity2)
	city.AddNeighbor("east", neighborCity3)
	city.AddNeighbor("west", neighborCity4)
	nn, _ := world.GetRandomNeighbor("Tokio", s)
	assert.NotEmpty(t, nn)
}

func TestGetRandomNeighborFail(t *testing.T) {
	var s int64
	test := "Tokio"
	city := NewCity(test)
	world := NewWorld()
	world.AddCity(&city)
	_, err := world.GetRandomNeighbor("Tokio", s)
	assert.Error(t, err)
}

func TestGetRandomNeighborEmptyWorld(t *testing.T) {
	var s int64
	test := "Tokio"
	city := NewCity(test)
	world := NewWorld()
	world.AddCity(&city)
	_, err := world.GetRandomNeighbor("Berlin", s)
	assert.Error(t, err)
}

//This is stale, can be removed
func BenchmarkNewCity(b *testing.B) {
	test := "London"
	for i := 0; i < b.N; i++ {
		_ = NewCity(test)
	}
}
