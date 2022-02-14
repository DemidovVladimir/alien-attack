package alien

type WorldUseCase interface {
	ProvideRandomCity(int64) (string, error)
}
