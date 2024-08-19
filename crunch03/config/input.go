package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadInput() [][]rune {
	// terminal input
	fmt.Println("Enter the height and width of your map (e.g., 1 2):")

	var h, w int
	_, err := fmt.Scanf("%d %d\n", &h, &w)

	for err != nil {
		fmt.Println("Invalid input. Please enter again:")
		var garbage string
		fmt.Scanf("%s\n", &garbage)
		_, err = fmt.Scanf("%d %d\n", &h, &w)
	}
	if h < 3 || w < 3 {
		fmt.Println(RED + "Invalid size value. Minimum grid size 3x3" + RESET)
		os.Exit(0)
	} else if h > 1000 || w > 1000 {
		fmt.Println(RED + "Invalid size value. Minimum grid size 10000x10000" + RESET)
		os.Exit(0)
	}
	arr := make([][]rune, h)
	for i := range arr {
		arr[i] = make([]rune, w)
	}
	// input rune map of cells
	for {
		fmt.Println("Enter the map row by row (use # for live cells and . for dead cells):")
		strArr := make([]string, h)
		valid := true

		for i := 0; i < h; i++ {
			_, err := fmt.Scanf("%s\n", &strArr[i])
			if len(strArr[i]) != w || err != nil {
				valid = false
				break
			}
			for _, char := range strArr[i] {
				if char != '#' && char != '.' {
					valid = false
					break
				}
			}
		}

		if valid {
			for i := 0; i < h; i++ {
				for j := 0; j < w; j++ {
					arr[i][j] = rune(strArr[i][j])
				}
			}
			break
		}
		fmt.Println("Invalid map input. Please try again.")
	}
	return arr
}

func ReadGridFromFile() ([][]rune, error) {
	// input from file
	data, err := os.ReadFile(FilePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	if len(lines) == 0 {
		return nil, errors.New("file is empty")
	}

	// Initialize the grid
	grid := make([][]rune, len(lines))

	// Validate line lengths
	expectedLength := len([]rune(lines[0]))
	for i, line := range lines {
		if len([]rune(line)) != expectedLength {
			return nil, errors.New("inconsistent line lengths in the file")
		}
		grid[i] = []rune(line)
	}
	return grid, nil
}
