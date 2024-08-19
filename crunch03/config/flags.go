package config

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	IsHelp        bool
	DelayMs       int
	IsVerbose     bool
	FilePath      string
	IsEdgePortals bool
	IsFullscreen  bool
	IsFootprint   bool
	IsColored     bool
	Width         int
	Height        int
	IsRandom      bool
	WinHeight     int
	WinWidth      int
)

func ParseFlags() {
	// initializer for all the flags
	randomSize := ""
	flag.Usage = func() {
		PrintHelp()
		os.Exit(0)
	}
	flag.BoolVar(&IsHelp, "help", false, "Print help message")
	flag.BoolVar(&IsVerbose, "verbose", false, "Enable verbose output")
	flag.IntVar(&DelayMs, "delay-ms", 2500, "Set delay in milliseconds")
	flag.StringVar(&FilePath, "file", "", "Path to input file")
	flag.BoolVar(&IsEdgePortals, "edges-portal", false, "Enable edge portals")
	flag.StringVar(&randomSize, "random", "", "Generate random map with dimensions WxH")
	flag.BoolVar(&IsFullscreen, "fullscreen", false, "Enable fullscreen mode")
	flag.BoolVar(&IsFootprint, "footprints", false, "Show cell footprints")
	flag.BoolVar(&IsColored, "colored", false, "Enable colored output")
	flag.Parse()

	// Fix order of execution
	if randomSize != "" && FilePath != "" {
		for _, v := range os.Args {

			if v[:8] == "--random" {
				FilePath = ""
				break
			} else if v[:6] == "--file" {
				randomSize = ""
				break
			}
		}
	}

	nonFlagArgs := flag.Args()
	if len(nonFlagArgs) > 0 {
		fmt.Println(RED + "Flag is not supported")
		os.Exit(0)
	}
	// validation
	if randomSize != "" {
		IsRandom = true
		dimensions := strings.Split(randomSize, "x")
		if len(dimensions) != 2 {
			fmt.Println(RED + "Invalid random size format. Use WxH.")
			os.Exit(0)
		}
		var err error
		Width, err = strconv.Atoi(dimensions[0])
		if err != nil {
			fmt.Println(RED + "Invalid width value.")
			os.Exit(0)
		}
		Height, err = strconv.Atoi(dimensions[1])
		if err != nil {
			fmt.Println(RED + "Invalid height value.")
			os.Exit(0)
		}
	}
	if IsFullscreen {
		WinWidth, WinHeight, _ = getTerminalSize()
	}
}

// gets parameters needed for fullscreen
func getTerminalSize() (int, int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}

	var rows, cols int
	if _, err := fmt.Sscanf(string(out), "%d %d", &rows, &cols); err != nil {
		return 0, 0, err
	}

	return cols, rows, nil
}

func PrintHelp() {
	fmt.Println("Usage: go run main.go [options]\n")
	fmt.Println("Options:")
	fmt.Println("  --help        : Show the help message and exit")
	fmt.Println("  --verbose     : Display detailed information about the simulation, including grid size, number of ticks, speed, and map name")
	fmt.Println("  --delay-ms=X  : Set the animation speed in milliseconds. Default is 2500 milliseconds")
	fmt.Println("  --file=X      : Load the initial grid from a specified file")
	fmt.Println("  --edges-portal: Enable portal edges where cells that exit the grid appear on the opposite side")
	fmt.Println("  --random=WxH  : Generate a random grid of the specified width (W) and height (H)")
	fmt.Println("  --fullscreen  : Adjust the grid to fit the terminal size with empty cells")
	fmt.Println("  --footprints  : Add traces of visited cells, displayed as '∘'")
	fmt.Println("  --colored     : Add color to live cells and traces if footprints are enabled")
	os.Exit(0)
}
