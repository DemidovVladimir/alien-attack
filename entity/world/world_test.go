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

//This is stale, can be removed
func BenchmarkNewCity(b *testing.B) {
	test := "London"
	for i := 0; i < b.N; i++ {
		_ = NewCity(test)
	}
}
