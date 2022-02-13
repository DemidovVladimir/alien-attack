package alien

import (
	"testing"

	"github.com/VladimirDemidov/alien-attack/internal/fs"
	"github.com/stretchr/testify/assert"
)

func TestChooseLocation(t *testing.T) {
	w, _ := fs.ReadWorldFile("../../static/world.txt")
	a := NewAlien("Bryval")
	l := ChooseLocation(w, a)
	assert.NotNil(t, l)
	assert.NotNil(t, a.Location)
}

func TestChooseLocationWithAnotherAlien(t *testing.T) {
	w, _ := fs.ReadWorldFile("../../static/onecityworld.txt")
	a := NewAlien("Bryval")
	ChooseLocation(w, a)
	b := NewAlien("Zroth")
	l := ChooseLocation(w, b)
	assert.Nil(t, l)
}

func TestNewAlien(t *testing.T) {
	a := NewAlien("Iroverk")
	assert.Equal(t, "Iroverk", a.Name)
}

func TestBattle(t *testing.T) {
	w, _ := fs.ReadWorldFile("../../static/onecityworld.txt")
	a := NewAlien("Bryval")
	c := ChooseLocation(w, a)
	b := NewAlien("Zroth")
	ChooseLocation(w, b)
	a.Battle(c, w)
	assert.Equal(t, len(LandedAliens), 1)
	assert.Nil(t, w.Cities[0])
}

func TestMove(t *testing.T) {
	w, _ := fs.ReadWorldFile("../../static/world.txt")
	a := NewAlien("Bryval")
	c := ChooseLocation(w, a)
	a.Move(w)
	assert.NotNil(t, a.Location)
	assert.Equal(t, 0, len(c.Aliens))
}

// func TestDie(t *testing.T) {
// 	NewAlien("Bryval")
// 	assert.Equal(t, 1, len(LandedAliens))
// 	// assert.Equal(t, a.Name, LandedAliens["Bryval"].Name)
// 	// a.Die()
// 	// assert.Nil(t, LandedAliens["Bryval"])
// 	// assert.Equal(t, 0, len(LandedAliens))
// }
