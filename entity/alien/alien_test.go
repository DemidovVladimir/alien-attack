package alien

import (
	"testing"

	"github.com/VladimirDemidov/alien-attack/internal/fs"
	"github.com/stretchr/testify/assert"
)

func TestChooseLocation(t *testing.T) {
	var i int64
	w, _ := fs.ReadWorldFile("../../static/world.txt")
	a := NewAlien("Bryval")
	ChooseLocation(w, a, i)
	assert.NotNil(t, a.Location)
}

func TestChooseLocationWithoutCity(t *testing.T) {
	var i int64
	w, _ := fs.ReadWorldFile("../../static/onecityworld.txt")
	w.Cities = nil
	a := NewAlien("Bryval")
	err := ChooseLocation(w, a, i)
	assert.Error(t, err)
}

func TestNewAlien(t *testing.T) {
	a := NewAlien("Iroverk")
	assert.Equal(t, "Iroverk", a.Name)
}

// func TestMove(t *testing.T) {
// 	var i int64
// 	c := make(chan string)
// 	w, _ := fs.ReadWorldFile("../../static/world.txt")
// 	a := NewAlien("Bryval")

// 	ChooseLocation(w, a, i)
// 	a.Move(w, i, c)
// 	assert.NotNil(t, a.Location)
// }

// func TestMoveFail(t *testing.T) {
// 	var i int64
// 	c := make(chan string)
// 	w, _ := fs.ReadWorldFile("../../static/world.txt")
// 	a := NewAlien("Bryval")

// 	ChooseLocation(w, a, i)
// 	a.Location.Directions = nil
// 	err := a.Move(w, i, c)
// 	assert.Error(t, err)
// }
