package main

import (
	"conway/config"
	"conway/controllers"
	"conway/models"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	config.ParseFlags()
	if config.IsHelp {
		config.PrintHelp()
		return
	}

	var game models.Map

	switch {
	case config.FilePath != "":
		gameMap, err := config.ReadGridFromFile()
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		game = models.NewMapFromRunes(gameMap)
	case config.IsRandom:
		game = models.NewRandomMap(config.Width, config.Height)
		if (config.WinWidth < game.Width() || config.WinHeight < game.Height()) && config.IsFullscreen {
			fmt.Printf("Too big size, the current screen size: %dx%d\n", config.WinWidth, config.WinHeight)
			os.Exit(0)
		}
	default:
		gameMap := config.ReadInput()
		game = models.NewMapFromRunes(gameMap)
	}

	game.BakeCount()
	runGame(&game)
}

func runGame(game *models.Map) {
	tick := 1 // counter
	for {
		// clears terminal at each iteration
		cmd := exec.Command("clear")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Print(string(output))
		// pritn fi flag verbose
		if config.IsVerbose {
			fmt.Printf("Tick: %d\nGrid Size: %dx%d\nLive Cells: %d\nDelayMs: %dms\n\n", tick, config.Width, config.Height, game.LiveCount, config.DelayMs)
		}

		controllers.PrintMap(*game)
		controllers.UpdateMap(game)
		// delay for tick
		time.Sleep(time.Duration(config.DelayMs) * time.Millisecond)
		tick++
	}
}
