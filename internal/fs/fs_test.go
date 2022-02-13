package fs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadWorldFile(t *testing.T) {
	world, _ := ReadWorldFile("../../static/world.txt")
	assert.Equal(t, 25, len(world.Cities))
}

func TestReadWorldFileReadFail(t *testing.T) {
	_, err := ReadWorldFile("../../world.txt")
	assert.Error(t, err)
}

func TestReadAliensFile(t *testing.T) {
	aliens, _ := ReadAliensFile("../../static/aliens.txt")
	assert.Equal(t, 30, len(aliens))
}

func TestReadAliensFileFailure(t *testing.T) {
	_, err := ReadAliensFile("../../static/stars.txt")
	assert.Error(t, err)
}
