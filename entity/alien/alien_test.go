package alien

import (
	"errors"
	"testing"

	"github.com/VladimirDemidov/alien-attack/usecase/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestChooseLocation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockWorldUseCase(ctrl)

	m.
		EXPECT().
		ProvideRandomCity(gomock.Eq(int64(0))).
		Return("test", nil)

	var i int
	a := NewAlien("Bryval")
	ChooseLocation(m, a, i)
	assert.NotNil(t, a.Location)
}

func TestChooseLocationWithoutCity(t *testing.T) {
	var i int
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockWorldUseCase(ctrl)

	m.
		EXPECT().
		ProvideRandomCity(gomock.Eq(int64(0))).
		Return("", errors.New("Error"))

	a := NewAlien("Bryval")
	_, err := ChooseLocation(m, a, i)
	assert.Error(t, err)
}

func TestNewAlien(t *testing.T) {
	a := NewAlien("Iroverk")
	assert.Equal(t, "Iroverk", a.Name)
}

func TestNewSwarm(t *testing.T) {
	a := NewSwarm()
	assert.NotNil(t, a)
}

func TestMove(t *testing.T) {
	var i int
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockWorldUseCase(ctrl)

	m.
		EXPECT().
		GetRandomNeighbor(gomock.Eq("test"), gomock.Eq(int64(0))).
		Return("test", nil)

	m.
		EXPECT().
		ProvideRandomCity(gomock.Eq(int64(0))).
		Return("test", nil)

	a := NewAlien("Bryval")
	al, _ := ChooseLocation(m, a, i)
	a.Location = al
	newLocation, _ := a.Move(m, i)
	assert.NotNil(t, newLocation)
}
