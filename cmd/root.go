/*
Copyright © 2022 Vladimir Demidov uncojet@gmail.com

*/
package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/spf13/cobra"
)

//Using cobra here, but I guess it is too much for this app
//it could be used in future to make an app more flexible
//For example I was thinking of adding config for different alien nations,
//here I am only using characters from Starcraft(Zergs and Protos), could be aliens
//from X-Com, Alienation or other game/movie
var (
	s  int
	w  string
	a  string
	wg sync.WaitGroup
	//Quit is a gracefull shutdown, would be good to have in case of service runtime
	quit    = make(chan os.Signal)
	rootCmd = &cobra.Command{
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

}
