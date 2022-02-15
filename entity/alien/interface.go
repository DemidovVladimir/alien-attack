package alien

type WorldUseCase interface {
	ProvideRandomCity(int64) (string, error)
	GetRandomNeighbor(string, int64) (string, error)
	GetCityByName(string) (interface{}, error)
}

type CityUseCase interface {
	AddAlienOrFight(name string) error
}

type AlienUseCase interface {
	Move(WorldUseCase, int) (string, error)
}
