package alien

import (
	"testing"

	"github.com/VladimirDemidov/alien-attack/internal/fs"
	"github.com/stretchr/testify/assert"
)

func TestChooseLocation(t *testing.T) {
	var i int64
	c := make(chan *Alien)
	w, _ := fs.ReadWorldFile("../../static/world.txt")
	a := NewAlien("Bryval", c)
	l, _ := ChooseLocation(w, a, i)
	assert.NotNil(t, l)
	assert.NotNil(t, a.Location)
}

func TestChooseLocationWithoutCity(t *testing.T) {
	var i int64
	c := make(chan *Alien)
	w, _ := fs.ReadWorldFile("../../static/onecityworld.txt")
	w.Cities = nil
	a := NewAlien("Bryval", c)
	_, err := ChooseLocation(w, a, i)
	assert.Error(t, err)
}

func TestNewAlien(t *testing.T) {
	c := make(chan *Alien)
	a := NewAlien("Iroverk", c)
	assert.Equal(t, "Iroverk", a.Name)
}

func TestMove(t *testing.T) {
	var i int64
	c := make(chan *Alien)
	w, _ := fs.ReadWorldFile("../../static/world.txt")
	a := NewAlien("Bryval", c)
	cc, _ := ChooseLocation(w, a, i)
	a.Move(w, i)
	assert.NotNil(t, a.Location)
	assert.Equal(t, 0, len(cc.Aliens))
}
