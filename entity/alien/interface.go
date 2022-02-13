package alien

import "github.com/VladimirDemidov/alien-attack/entity/world"

type UseCase interface {
	Move(w *world.World)
	Battle(c *world.City, w *world.World)
	Die()
}
