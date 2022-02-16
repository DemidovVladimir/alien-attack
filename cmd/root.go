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

	"github.com/VladimirDemidov/alien-attack/config"
	"github.com/VladimirDemidov/alien-attack/entity/alien"
	"github.com/VladimirDemidov/alien-attack/entity/world"
	"github.com/VladimirDemidov/alien-attack/internal/fs"
	"github.com/spf13/cobra"
)

//Using cobra here, but I guess it is too much for this app
//it could be used in future to make an app more flexible
//For example I was thinking of adding config for different alien nations,
//here I am only using characters from Starcraft(Zergs and Protos), could be aliens
//from X-Com, Alienation or other game/movie
var (
	s int
	w string
	a string
	//Quit is a gracefull shutdown, would be good to have in case of service runtime
	quit = make(chan os.Signal)
	//Remove alien channel
	rac = make(chan string, 2)
	//Remove city channel
	rcc     = make(chan string)
	wg      sync.WaitGroup
	rootCmd = &cobra.Command{
		Use:   "alien-attack",
		Short: "An alien attack game simulation",
		Long: `An alien attack game simulation, please pick how massaive the attack is 
		going to be. It will generate the world and move aliens arround, later on we can add
		alien specialities and abilities. This way it will be interesting to run the simulation.
		Other then that there could be an option to invade different other worlds, current version
		is only invading world of Japan`,
		Run: func(cmd *cobra.Command, args []string) {
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
			fmt.Println(`
	Welcome to the inspiring web3 universe, where aliens are not the biggest topic anymore...
	Here is the world and invaders configuration, upon which you would act on:
	`)

			fmt.Println("Invaders swarm flying to the planet (swarm consists of)", s, "creatures")
			fmt.Println("You have configured this word to be invaded (world config file)", w)
			fmt.Println(`Aliens could be pretty charming creatures, depends on your preference,
	not this time, universe are doomed (creatures config file):`, a)

			//Initiate unoverse
			www, swarm := initiateUniverse(w, a)

			//Land all aliens, attack can be performed even before all of them are landed
			for i, a := range swarm.LandedAliens {
				wg.Add(1)
				go func(a string, w *world.World, swarm *alien.Swarm, i int, rcc chan string, wg *sync.WaitGroup) {
					defer wg.Done()
					if selectedAlien, k := swarm.Aliens.Load(a); k {
						sa := selectedAlien.(*alien.Alien)
						cn, _ := alien.ChooseLocation(w, selectedAlien.(*alien.Alien), i)
						if city, ok := www.Cities.Load(cn); ok {
							cc := city.(*world.City)
							cc.Aliens = append(cc.Aliens, a)
							sa.Location = cn
						}

					}
				}(a, www, swarm, i, rcc, &wg)
			}

			//Performing aliens movement, with battles meanwhile
			for _, a := range swarm.LandedAliens {
				wg.Add(1)
				go func(
					a string,
					w *world.World,
					swarm *alien.Swarm,
					rcc chan string,
					rac chan string,
					wg *sync.WaitGroup,
				) {
					defer wg.Done()
					for i := 0; i < config.STEPS; i++ {
						if selectedAlien, k := swarm.Aliens.Load(a); k {
							sa := selectedAlien.(*alien.Alien)
							cn, _ := sa.Move(www, i)
							if city, ok := www.Cities.Load(cn); ok {
								cc := city.(*world.City)
								//If there is an alien already, kill it and current alien
								if len(cc.Aliens) > 0 {
									rac <- a
									rcc <- cn
								} else {
									cc.Aliens = append(cc.Aliens, a)
								}
							}
						}
					}
				}(a, www, swarm, rcc, rac, &wg)
			}

			go func(rac chan string, rcc chan string, swarm *alien.Swarm, www *world.World, wg *sync.WaitGroup) {
				for {
					select {
					case al := <-rac:
						swarm.Aliens.Delete(al)
					case ci := <-rcc:
						www.Cities.Delete(ci)
					case <-quit:
						return
					}
				}
			}(rac, rcc, swarm, www, &wg)

			wg.Wait()
			quit <- syscall.SIGQUIT
			fmt.Println(`
			End phase...
			World map after the alien attack:
			`)
			www.Cities.Range(func(key, value interface{}) bool {
				fmt.Println("City:", key)
				return true
			})
			fmt.Println(`
			End phase...
			Ready to be evacurated:
			`)
			swarm.Aliens.Range(func(key, value interface{}) bool {
				fmt.Println("Alien:", key)
				return true
			})
		},
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

//Read world file and aliens file
//Receive user input to read files
//Returns pointer to the world and swarm structs
func initiateUniverse(w, a string) (*world.World, *alien.Swarm) {
	www, err := fs.ReadWorldFile(w)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(`
			Starting phase...
			Map of the inocent world:
			`)
	www.Cities.Range(func(key, value interface{}) bool {
		fmt.Println("City able to fight back:", key)
		return true
	})

	//Get all available alien names
	swarm, err := fs.ReadAliensFile(a)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(`
			Starting phase...
			Wariors ready to begin:
			`)
	swarm.Aliens.Range(func(key, value interface{}) bool {
		fmt.Println("Invader", key)
		return true
	})
	return www, swarm
}
