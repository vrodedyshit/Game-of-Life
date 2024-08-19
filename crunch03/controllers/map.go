package controllers

import (
	"conway/config"
	"conway/models"
	"fmt"
	"os"
	"os/exec"
)

func PrintMap(gameMap models.Map) {
	live := '×'
	dead := '·'
	trace := '·'
	spacing := ""
	nextLine := "\n"
	var liveColor string = config.RESET
	var deadColor string = config.RESET
	var traceColor string = config.RESET
	// if flag colored
	if config.IsColored {
		liveColor = config.GREEN
		deadColor = config.RED
		traceColor = deadColor
	}
	// if flag footprints
	if config.IsFootprint {
		trace = '∘'
		if config.IsColored {
			traceColor = config.YELLOW
		}
	}
	// print case for fullscreen
	if config.IsFullscreen {

		for (2+len(spacing))*(gameMap.Width()-1) < config.WinWidth {
			spacing += " "
		}
		for (3+len(nextLine))*gameMap.Height() < config.WinHeight {
			nextLine += "\n"
		}
	} else {
		spacing = " "
	}
	// print the map
	for _, rowCell := range gameMap.Cells {
		for i, cell := range rowCell {
			if i != len(rowCell)-1 {
				if cell.Symbol == '#' {
					fmt.Printf("%s%c%s%s", liveColor, live, config.RESET, spacing)
				} else if cell.Symbol == '.' && !cell.IsFootprint {
					fmt.Printf("%s%c%s%s", deadColor, dead, config.RESET, spacing)
				} else {
					fmt.Printf("%s%c%s%s", traceColor, trace, config.RESET, spacing)
				}
			} else {
				if cell.Symbol == '#' {
					fmt.Printf("%s%c%s", liveColor, live, config.RESET)
				} else if cell.Symbol == '.' && !cell.IsFootprint {
					fmt.Printf("%s%c%s", deadColor, dead, config.RESET)
				} else {
					fmt.Printf("%s%c%s", traceColor, trace, config.RESET)
				}
			}

		}
		fmt.Printf("%s", nextLine)
	}
}

func UpdateMap(gameMap *models.Map) {
	gameMap.LiveCount = 0
	width := gameMap.Width()
	height := gameMap.Height()

	for _, rowCell := range gameMap.Cells {
		for _, cell := range rowCell {
			// underpopulation
			if cell.IsAlive && cell.NeighborCount < 2 {
				cell.IsAlive = false
				cell.IsFootprint = true
				cell.Symbol = '·'
			}
			// overpopulation
			if cell.IsAlive && cell.NeighborCount > 3 {
				cell.IsAlive = false
				cell.IsFootprint = true
				cell.Symbol = '·'
			}
			// reproduction
			if !cell.IsAlive && cell.NeighborCount == 3 {
				cell.IsAlive = true
				cell.Symbol = '#'
			}
			if cell.IsAlive {
				gameMap.LiveCount++
			}
			gameMap.UpdateCell(cell)
		}
	}

	if config.IsEdgePortals {
		for y, rowCell := range gameMap.Cells {
			for x, _ := range rowCell {
				cell := &gameMap.Cells[y][x]
				// Apply edge portal rules
				if cell.IsAlive {
					if y == 0 { // top edge
						gameMap.Cells[height-1][x].IsAlive = true
						gameMap.Cells[height-1][x].Symbol = '#'
					}
					if y == height-1 { // bottom edge
						gameMap.Cells[0][x].IsAlive = true
						gameMap.Cells[0][x].Symbol = '#'
					}
					if x == 0 { // left edge
						gameMap.Cells[y][width-1].IsAlive = true
						gameMap.Cells[y][width-1].Symbol = '#'
					}
					if x == width-1 { // right edge
						gameMap.Cells[y][0].IsAlive = true
						gameMap.Cells[y][0].Symbol = '#'
					}
				}
			}
		}
	}
	// stopping the game if no live cells left
	if gameMap.LiveCount <= 0 {
		cmd := exec.Command("clear")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Print(string(output))
		PrintMap(*gameMap)
		fmt.Println(config.RED + "no living cells left")
		os.Exit(0)
	}
	gameMap.BakeCount()
}
