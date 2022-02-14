package world

type WorldUseCase interface {
	AddCity(*City)
	RemoveCity(*City)
	ProvideRandomCity(int64) (string, error)
}

type CityUseCase interface {
	AddNeighbor(int, *City) error
}
