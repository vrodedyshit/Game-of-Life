package models

type Offset [][]int

var (
	Neighbours Offset = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	Corners    Offset = [][]int{{1, 1}, {-1, 1}, {1, -1}, {-1, -1}}
)

type Map struct {
	width     int
	height    int
	LiveCount int
	Cells     [][]Cell
}

func (gameMap Map) Width() int {
	return gameMap.width
}

func (gameMap Map) Height() int {
	return gameMap.height
}

func (gameMap Map) BakeCount() {
	for _, row := range gameMap.Cells {
		for _, cell := range row {
			neighbours := gameMap.GetNeighbours(cell, Neighbours, Corners)

			mineCount := 0
			for _, neighbour := range neighbours {
				if !neighbour.IsAlive {
					continue
				}
				mineCount++
			}
			cell.NeighborCount = mineCount
			gameMap.UpdateCell(cell)
		}
	}
}

func (gameMap Map) GetNeighbours(cell Cell, offsets ...Offset) []Cell {
	result := make([]Cell, 0, 8)
	for _, offset := range offsets {
		for _, coordinates := range offset {
			offsetX := coordinates[0]
			offsetY := coordinates[1]
			neighbourCell, found := gameMap.GetCell(cell.x+offsetX, cell.y+offsetY)
			if !found {
				continue
			}
			result = append(result, neighbourCell)
		}
	}
	return result
}

func (gameMap Map) GetCell(x, y int) (Cell, bool) {
	if x < 0 || y < 0 {
		return Cell{}, false
	}
	if x >= gameMap.Width() || y >= gameMap.Height() {
		return Cell{}, false
	}
	return gameMap.Cells[y][x], true
}

func (gameMap Map) UpdateCell(cell Cell) {
	gameMap.Cells[cell.y][cell.x] = cell
}
