package controller

import (
	"fmt"
	"os"

	"github.com/Arafatk/glot"
)

func test() {
	dimensions := 2
	// The dimensions supported by the plot
	persist := true
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)
	pointGroupName := "Simple Circles"
	style := "lines"
	points := []float64{12, 13, 11, 1, 7}
	// Adding a point group
	plot.AddPointGroup(pointGroupName, style, points)
	// pointGroupName = "Simple Lines"
	// style = "lines"
	// points = [][]float64{{7, 3, 3, 5.6, 5.6, 7, 7, 9, 13, 13, 9, 9}, {10, 10, 4, 4, 5.4, 5.4, 4, 4, 4, 10, 10, 4}}
	// plot.AddPointGroup(pointGroupName, style, points)
	// A plot type used to make points/ curves and customize and save them as an image.
	plot.SetTitle("Example Plot")
	// Optional: Setting the title of the plot
	plot.SetXLabel("X-Axis")
	plot.SetYLabel("Y-Axis")
	// Optional: Setting label for X and Y axis
	plot.SetXrange(-2, 18)
	plot.SetYrange(-2, 18)
	// Optional: Setting axis ranges
	plot.SavePlot("2.png")
}

// NewDegreeDistributionPlot displays and saves a graph of the plot and optionally a csv
func NewDegreeDistributionPlot(networkName string, adjacencyMatrix [][]uint8, saveCSV bool) {
	degreeDistribution := makeDegreeDistribution(adjacencyMatrix)
	plotDegreeDistribution(networkName, degreeDistribution)
	if saveCSV {
		writeCSV(networkName, degreeDistribution)
	}
}

func writeCSV(networkName string, degreeDistribution []uint32) {
	csv, err := os.Create(networkName + ".csv")
	if err != nil {
		panic(err)
	}
	defer csv.Close()

	fmt.Fprintln(csv, "Degree Distribution")
	for degree, frequency := range degreeDistribution {
		fmt.Fprintf(csv, "%d,%d\n", degree, frequency)
	}
}

func plotDegreeDistribution(networkName string, degreeDistribution []uint32) {
	dimensions := 2
	persist := true
	debug := false
	plot, err := glot.NewPlot(dimensions, persist, debug)
	if err != nil {
		panic(err)
	}
	pointGroupName := networkName
	style := "lines"
	plot.AddPointGroup(pointGroupName, style, u32tof64(degreeDistribution))
	plot.SetTitle(networkName)
	plot.SetXLabel("Degree")
	plot.SetYLabel("Frequency")
	plot.SetXrange(0, len(degreeDistribution))
	maxFrequency := uint32(0)
	for _, frequency := range degreeDistribution {
		if frequency > maxFrequency {
			maxFrequency = frequency
		}
	}
	plot.SetYrange(0, int(maxFrequency))
	if err = plot.SavePlot(networkName + ".png"); err != nil {
		panic(err)
	}
}

func u32tof64(slice []uint32) []float64 {
	converted := make([]float64, len(slice))
	for i, val := range slice {
		converted[i] = float64(val)
	}
	return converted
}

func makeDegreeDistribution(adjacencyMatrix [][]uint8) []uint32 {
	nodeDegrees := make([]uint32, len(adjacencyMatrix))
	count := uint32(0)
	for i, row := range adjacencyMatrix {
		for _, edge := range row {
			if edge > 0 {
				count++
			}
		}
		nodeDegrees[i] = count
		count = 0
	}

	degreeCounter := make(map[uint32]uint32)
	for _, degree := range nodeDegrees {
		degreeCounter[degree]++
	}

	maxDegree := uint32(0)
	for degree := range degreeCounter {
		if degree > maxDegree {
			maxDegree = degree
		}
	}
	degreeDistribution := make([]uint32, maxDegree)
	for degree, cnt := range degreeCounter {
		degreeDistribution[degree-1] = cnt
	}

	return degreeDistribution
}
