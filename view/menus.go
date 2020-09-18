package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/GaudiestTooth17/social_circles_go/model"

	"github.com/GaudiestTooth17/social_circles_go/controller"
)

// MainMenu is the main entry point of the program
func MainMenu(configuration controller.NetworkConfig) {
	fmt.Println("0) exit\n1) generate network\n2) edit configuration")
	choice := readLn()
	if choice == "0" {
		return
	} else if choice == "1" {
		generateNetworkMenu(configuration, make([][]uint8, 0))
		MainMenu(configuration)
	} else if choice == "2" {
		editMenu(configuration)
		MainMenu(configuration)
	} else {
		fmt.Print("\nEnter a valid choice.\n\n")
		MainMenu(configuration)
	}
}

func generateNetworkMenu(configuration controller.NetworkConfig, adjMat [][]uint8) {
	fmt.Println("0) return to main menu\n1) generate and view plot\n2) save matrix")
	choice := readLn()
	if choice == "0" {
		MainMenu(configuration)
	} else if choice == "1" {
		adjMat = controller.NewAdjacencyMatrix(configuration)
		controller.NewDegreeDistributionPlot(configuration.NetworkName,
			adjMat, configuration.SaveCSV)
		generateNetworkMenu(configuration, adjMat)
	} else if choice == "2" {
		controller.SaveAdjacencyList(configuration.NetworkName, adjMat)
		generateNetworkMenu(configuration, adjMat)
	} else {
		fmt.Print("\nEnter a valid choice.\n\n")
		generateNetworkMenu(configuration, adjMat)
	}
}

func editMenu(config controller.NetworkConfig) {
	fmt.Println(`0) return to main menu
1) set network name
2) set NoIsolates
3) set SaveCSV
4) set grid size
5) add/edit reach configuration
6) remove reach configuration`)
	choice := readLn()

	switch choice {
	case "0":
		MainMenu(config)
	case "1":
		fmt.Print("Enter name: ")
		name := readLn()
		config.NetworkName = name
		editMenu(config)
	case "2":
		fmt.Printf("NoIsolates is currently %v.\n", config.NoIsolates)
		fmt.Print("Enter t or f ")
		noIsolates := readLn()
		if noIsolates == "t" {
			config.NoIsolates = true
		} else {
			config.NoIsolates = false
		}
		editMenu(config)
	case "3":
		fmt.Printf("SaveCSV is currently %v.\n", config.SaveCSV)
		fmt.Print("Enter t or f ")
		saveCSV := readLn()
		if saveCSV == "t" {
			config.SaveCSV = true
		} else {
			config.SaveCSV = false
		}
		editMenu(config)
	case "4":
		fmt.Printf("Grid size is currently %d.\n", config.GridSize)
		fmt.Print("New size: ")
		sizeStr := readLn()
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			panic(err)
		}
		config.GridSize = int32(size)
		editMenu(config)
	case "5":
		addReachConfig(&config)
		editMenu(config)
	case "6":
		removeReachConfig(&config)
		editMenu(config)
	}
}

// addReachConfig can be used to either add or edit an existing ReachConfig
func addReachConfig(config *controller.NetworkConfig) {
	stdinReader := bufio.NewReader(os.Stdin)
	config.PrintReaches()

	fmt.Print("Agent type name: ")
	name, _ := stdinReader.ReadString('\n')
	name = name[:len(name)-1]

	fmt.Print("Agent quantity: ")
	qStr, _ := stdinReader.ReadString('\n')
	quantity, _ := strconv.Atoi(qStr[:len(qStr)-1])

	fmt.Print("Agent reach: ")
	reachStr, _ := stdinReader.ReadString('\n')
	reach, _ := strconv.Atoi(reachStr[:len(reachStr)-1])

	config.ReachConfigs[name] = model.ReachConfig{Quantity: uint32(quantity),
		Reach: uint8(reach)}
}

func removeReachConfig(config *controller.NetworkConfig) {
	stdinReader := bufio.NewReader(os.Stdin)
	config.PrintReaches()

	fmt.Print("Which reach will be deleted? ")
	name, _ := stdinReader.ReadString('\n')
	name = name[:len(name)-1]
	if _, ok := config.ReachConfigs[name]; ok {
		delete(config.ReachConfigs, name)
	}
}

func makeAndPlotAdjMat(config controller.NetworkConfig) {
	adjacencyMatrix := controller.NewAdjacencyMatrix(config)
	controller.NewDegreeDistributionPlot(config.NetworkName, adjacencyMatrix,
		config.SaveCSV)
}

func readLn() string {
	stdinReader := bufio.NewReader(os.Stdin)
	str, _ := stdinReader.ReadString('\n')
	return str[:len(str)-1]
}
