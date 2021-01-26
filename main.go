package main

import (
	"github.com/GaudiestTooth17/social_circles_go/controller"
	"github.com/GaudiestTooth17/social_circles_go/model"
)

func main() {
	// emptyConfiguration := controller.NetworkConfig{ReachConfigs: make(map[string]model.ReachConfig),
	// 	NoIsolates: false, GridSize: 0, SaveCSV: false}
	// view.MainMenu(emptyConfiguration)
	configuration := controller.NetworkConfig{
		ReachConfigs: map[string]model.ReachConfig{
			"green": model.ReachConfig{Quantity: 500, Reach: 5},
		},
		GridSize:    65,
		NoIsolates:  false,
		SaveCSV:     false,
		NetworkName: "spatial-network",
	}
	graph := controller.NewAdjacencyMatrix(configuration)
	controller.SaveAdjacencyList(configuration.NetworkName, graph, true)
}
