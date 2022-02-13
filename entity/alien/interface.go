package alien

import "github.com/VladimirDemidov/alien-attack/entity/world"

type UseCase interface {
	Move(*world.World, int64)
}
