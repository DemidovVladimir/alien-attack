/*
Copyright © 2022 Vladimir Demidov uncojet@gmail.com

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/VladimirDemidov/alien-attack/entity/alien"
	"github.com/VladimirDemidov/alien-attack/entity/world"
	"github.com/VladimirDemidov/alien-attack/internal/fs"
	"github.com/spf13/cobra"
)

type LiveAliens struct {
	m map[string]*alien.Alien
	sync.RWMutex
}

//Using cobra here, but I guess it is too much for this app
//it could be used in future to make an app more flexible
//For example I was thinking of adding config for different alien nations,
//here I am only using characters from Starcraft(Zergs and Protos), could be aliens
//from X-Com, Alienation or other game/movie
var (
	s          int
	w          string
	a          string
	liveAliens = LiveAliens{
		m: make(map[string]*alien.Alien),
	}
	listLanded []string
	//Quit is a gracefull shutdown, would be good to have in case of service runtime
	quit      = make(chan os.Signal)
	eliminate = make(chan string)
	rootCmd   = &cobra.Command{
		Use:   "alien-attack",
		Short: "An alien attack game simulation",
		Long: `An alien attack game simulation, please pick how massaive the attack is 
		going to be. It will generate the world and move aliens arround, later on we can add
		alien specialities and abilities. This way it will be interesting to run the simulation.
		Other then that there could be an option to invade different other worlds, current version
		is only invading world of Japan`,
		Run: RunRoot,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&s, "swarm", "s", 10, "How big is the invadors swarm")
	rootCmd.Flags().StringVarP(&w, "world", "w", "world.txt", "File describing world to invade")
	//This part can be replaced with config tag, not in this version
	rootCmd.Flags().StringVarP(&a, "aliens", "a", "aliens.txt", "File with alien names")
}

func RunRoot(cmd *cobra.Command, args []string) {
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	fmt.Println(`
	Welcome to the inspiring web3 universe, where aliens are not the biggest topic anymore...
	Here is the world and invaders configuration, upon which you would act on:
	`)

	fmt.Println("Invaders swarm flying to the planet (swarm consists of)", s, "creatures")
	fmt.Println("You have configured this word to be invaded (world config file)", w)
	fmt.Println(`Aliens could be pretty charming creatures, depends on your preference, 
	not this time, universe are doomed (creatures config file):`, a)

	www, alienSwarm := initiateUniverse(w, a)

	leftAlive := landing(www, alienSwarm, &liveAliens)

	if len(leftAlive.m) == 0 {
		fmt.Println("All aliens are dead, nice try")
		return
	}

	leftVeterans, www := invade(leftAlive, www)

	fmt.Println(len(leftVeterans.m), " live aliens after invasion, probably trapped")
	fmt.Println(len(www.Cities), " cities left able to fight back the swarm")
}

func initiateUniverse(w, a string) (*world.World, []string) {
	www, err := fs.ReadWorldFile(w)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Inocent world consist of:", len(www.Cities), "cities")

	//Get all available alien names
	allAlienNames, err := fs.ReadAliensFile(a)
	if err != nil {
		log.Fatal(err.Error())
	}

	//Land only those that fits into our alien swarm
	alienSwarm := allAlienNames[0:s]

	return www, alienSwarm
}

func landing(www *world.World, alienSwarm []string, liveAliens *LiveAliens) *LiveAliens {
	//Landing process
	//Land preconfigured swarm
	for i, alienName := range alienSwarm {
		//Here we got landed invaders
		la := alien.NewAlien(alienName)
		//Aliens are mostly invading big cities, city where alien landed
		cityName, _ := alien.ChooseLocation(www, la, int64(i))
		city, _ := www.GetCityByName(cityName)
		if address, ok := city.(*world.City); ok {
			err := address.AddAlienOrFight(la.Name)
			if err != nil {
				www.RemoveCity(address)
				for _, invader := range address.Aliens {
					delete(liveAliens.m, invader)
				}
			}
			liveAliens.m[la.Name] = la
			la.Location = address.Name
		}
	}

	return liveAliens
}

func invade(liveAliens *LiveAliens, www *world.World) (*LiveAliens, *world.World) {
	for k, ali := range liveAliens.m {
		for j := 0; j < 10000; j++ {
			if inv, ok := liveAliens.m[k]; ok {
				//Returns name of the next move if available
				nextMove, _ := inv.Move(www, int64(j))
				//Take battle in case of another alien in place
				city, _ := www.GetCityByName(nextMove)

				if address, ok := city.(*world.City); ok {
					err := address.AddAlienOrFight(ali.Name)
					if err != nil {
						www.RemoveCity(address)
						for _, invader := range address.Aliens {
							delete(liveAliens.m, invader)
						}
					}
					ali.Location = address.Name
				}
			}
		}
	}
	return liveAliens, www
}
