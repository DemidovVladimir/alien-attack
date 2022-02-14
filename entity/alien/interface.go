package alien

type WorldUseCase interface {
	ProvideRandomCity(int64) (string, error)
	GetCityByName(string) (interface{}, error)
	GetRandomNeighbor(string, int64) (string, error)
}

type CityUseCase interface {
	AddAlienOrFight(name string) error
}

type AlienUseCase interface {
	Move(WorldUseCase, int64) (string, error)
}
