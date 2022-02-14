package world

type WorldUseCase interface {
	AddCity(*City)
	RemoveCity(*City)
	ProvideRandomCity(int64) (string, error)
	GetCityByName(string) (interface{}, error)
	GetRandomNeighbor(string, int64) (string, error)
}

type CityUseCase interface {
	AddNeighbor(int, *City) error
	AddAlienOrFight(name string) error
}
