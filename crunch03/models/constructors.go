package models

import (
	"conway/config"
	"fmt"
	"math/rand"
	"os"
)

func newMap(cells [][]Cell, width int, height int) *Map {
	result := new(Map)

	result.Cells = cells
	result.width = width
	result.height = height

	return result
}

func NewCell(x, y int, symbol rune) *Cell {
	cell := new(Cell)

	cell.x = x
	cell.y = y
	cell.IsAlive = true
	cell.IsFootprint = false
	cell.Symbol = symbol
	if cell.Symbol == '#' {

		cell.IsAlive = true
	} else if cell.Symbol == '.' {
		cell.IsAlive = false
	} else {
		cell.IsAlive = false
		fmt.Println(config.RED + "Unsupported format! Use # for live cells and . for dead cells")
		os.Exit(0)
	}
	return cell
}

func NewRandomMap(width, height int) Map {
	gameMap := make([][]rune, height)

	var liveCount = 4

	liveCount += (width * height) / 10
	// Fill empty cells
	for y := 0; y < height; y++ {
		gameMap[y] = make([]rune, width)
		for x := 0; x < width; x++ {
			gameMap[y][x] = '.'
		}
	}

	// Put Live Cells to a random coordinates
	for cell := 0; cell < liveCount; cell++ {
		x := rand.Intn(width)
		y := rand.Intn(height)
		if gameMap[y][x] == '#' {
			cell--
		} else if gameMap[y][x] == '.' {
			gameMap[y][x] = '#'
		}

	}

	return NewMapFromRunes(gameMap)
}

func NewMapFromRunes(runes [][]rune) Map {
	cells := make([][]Cell, len(runes))
	for row, _ := range cells {
		cells[row] = make([]Cell, len(runes[0]))
	}
	height := len(runes)
	width := len(runes[0])
	gameMap := newMap(cells, width, height)
	gameMap.LiveCount = 0
	for row, _ := range runes {
		for col, element := range runes[row] {

			cell := NewCell(col, row, element)
			gameMap.Cells[row][col] = *cell
			if gameMap.Cells[row][col].IsAlive {
				gameMap.LiveCount++
			}
		}
	}

	return *gameMap
}
