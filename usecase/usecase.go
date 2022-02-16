package usecase

type WorldUseCase interface {
	ProvideRandomCity(int64) (string, error)
	GetRandomNeighbor(string, int64) (string, error)
	GetCityByName(string) (interface{}, error)
}
