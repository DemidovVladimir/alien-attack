package fs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadWorldFile(t *testing.T) {
	world, _ := ReadWorldFile("../../static/world.txt")
	_, ok := world.Cities.Load("Yokohama")
	assert.True(t, ok)
}

func TestReadWorldFileReadFail(t *testing.T) {
	_, err := ReadWorldFile("../../world.txt")
	assert.Error(t, err)
}

func TestReadAliensFile(t *testing.T) {
	aliens, _ := ReadAliensFile("../../static/aliens.txt")
	_, ok := aliens.Aliens.Load("Kikk")
	assert.True(t, ok)
}

func TestReadAliensFileFailure(t *testing.T) {
	_, err := ReadAliensFile("../../static/stars.txt")
	assert.Error(t, err)
}
