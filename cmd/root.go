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

	world, err := fs.ReadWorldFile(w)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("World consist of:", len(world.Cities), "cities")

	//Get all available alien names
	allAlienNames, err := fs.ReadAliensFile(a)
	if err != nil {
		log.Fatal(err.Error())
	}

	//Land only those that fits into our alien swarm
	alienSwarm := allAlienNames[0:s]

	for i, name := range alienSwarm {
		//Here we got landed invaders
		la := alien.NewAlien(name)
		//Aliens are mostly invading big cities, city where alien landed
		err := alien.ChooseLocation(world, la, int64(i))

		fmt.Println(err)

		// listLanded = append(listLanded, la.Name)
		// if err != nil {
		// 	log.Println(err.Error())
		// }
		// fmt.Println(la.Name, " has landed in", la.Location.Name)
		// liveAliens.m[name] = la
	}

	//Land aliens and start moving with go routines, as a realtime startegy
	//instead of step by step approach, but this is still not the final version
	// go func(c chan string, ll []string, la *LiveAliens) {
	// 	for i, alienStart := range listLanded {
	// 		for j := 0; j < 10000; j++ {
	// 			if _, ok := liveAliens.m[alienStart]; ok {
	// 				liveAliens.m[alienStart].Move(world, int64(i), c)
	// 			}
	// 		}
	// 	}
	// 	quit <- syscall.SIGQUIT
	// }(eliminate, listLanded, &liveAliens)

	// // Remove aliens from live slice
	// go func(c chan string) {
	// 	for name := range c {
	// 		delete(liveAliens.m, name)
	// 	}
	// }(eliminate)

	// <-quit

	fmt.Println(len(liveAliens.m), "aliens left alive after 10000 steps been performed...")
}
