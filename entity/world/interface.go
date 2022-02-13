package world

type WorldUseCase interface {
	AddCity(*City)
	RemoveCity(*City)
	GenerateWorld() error
}

type CityUseCase interface {
	AddNeighbor(int, *City) error
	NeighborStillExists(d int, w *World) bool
}
