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
	world.AddCity(city)
	assert.Equal(t, len(world.Cities), 1)
}

func TestRemoveCity(t *testing.T) {
	test := "Tokio"
	neighbor := "Minato"
	city := NewCity(test)
	neighborCity := NewCity(neighbor)
	world := NewWorld()
	world.AddCity(city)
	world.AddCity(neighborCity)
	world.RemoveCity(city)
	assert.Equal(t, neighborCity, world.AvailableCities[1])
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
	city.AddNeighbor(1, neighborCity)
	assert.Equal(t, city.Directions[1].Name, neighbor)
}

func TestNeighborStillExists(t *testing.T) {
	test := "Tokio"
	neighbor := "Minato"
	city := NewCity(test)
	neighborCity := NewCity(neighbor)
	world := NewWorld()
	world.AddCity(city)
	world.AddCity(neighborCity)
	city.AddNeighbor(1, neighborCity)
	//Validate if we have a neighbor
	world.RemoveCity(neighborCity)
	//Validate if our neighbor was defeated
	assert.Equal(t, 2, len(world.AvailableCities))
}

func TestAddNeighborFailure(t *testing.T) {
	test := "Tokio"
	neighbor1 := "Minato"
	neighbor2 := "Shinagava"
	city := NewCity(test)
	neighborCity1 := NewCity(neighbor1)
	neighborCity2 := NewCity(neighbor2)
	city.AddNeighbor(1, neighborCity1)
	err := city.AddNeighbor(1, neighborCity2)
	assert.Equal(t, err, errors.New("there is a neighbor at this direction"))
}

//This is stale, can be removed
func BenchmarkNewCity(b *testing.B) {
	test := "London"
	for i := 0; i < b.N; i++ {
		_ = NewCity(test)
	}
}
