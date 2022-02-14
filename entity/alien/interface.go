package alien

type WorldUseCase interface {
	ProvideRandomCity(int64) (string, error)
}

type CityUseCase interface {
	AddAlienOrFight(name string) error
}
