package main

import (
	"github.com/GaudiestTooth17/social_circles_go/controller"
	"github.com/GaudiestTooth17/social_circles_go/model"
	"github.com/GaudiestTooth17/social_circles_go/view"
)

func main() {
	emptyConfiguration := controller.NetworkConfig{ReachConfigs: make(map[string]model.ReachConfig),
		NoIsolates: false, GridSize: 0, SaveCSV: false}
	view.MainMenu(emptyConfiguration)
}
