package fs

import (
	"bufio"
	"os"
	"strings"

	"github.com/VladimirDemidov/alien-attack/entity/world"
)

//Read world file and generate world, not perfect, some side effects are happening
func ReadWorldFile(l string) (*world.World, error) {
	newWorld := world.NewWorld()
	f, err := os.Open(l)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		chunks := strings.Split(text, " ")
		city := world.NewCity(chunks[0])
		directions := chunks[1:]
		for j := 0; j < len(directions); j++ {
			ds := strings.Split(directions[j], "=")
			city.AddNeighbor(ds[0], world.NewCity(ds[1]))
		}
		newWorld.AddCity(&city)
	}

	return &newWorld, nil
}

//Read aliens file and generate world, not perfect, some side effects are happening
func ReadAliensFile(l string) (aliens []string, err error) {
	f, err := os.Open(l)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		alien := scanner.Text()
		aliens = append(aliens, alien)
	}
	return
}
