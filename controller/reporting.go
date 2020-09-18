package controller

// NetworkStats holds important statistical information for a social network
type NetworkStats struct {
	MeanDegree   float64
	MedianDegree float64
	ModeDegree   uint32
	MaxDegree    uint32
	MinDegree    uint32
}

// MakeStats for the provided social network
// func MakeStats(degreeDistribution []uint32) NetworkStats {
// 	//
// }

func calcMean(degreeDistribution []uint32) float64 {
	numNodes := uint32(0)
	for _, frequency := range degreeDistribution {
		numNodes += frequency
	}

	mean := float64(0)
	for i, frequency := range degreeDistribution {
		mean += (float64(i) / float64(numNodes)) * float64(frequency)
	}
	return mean
}

func calcMedian(degreeDistribution []uint32) float64 {
	numNodes := uint32(0)
	for _, frequency := range degreeDistribution {
		numNodes += frequency
	}

	nodeIndex := uint32(0)
	i := 0
	for nodeIndex < numNodes/2 {
		nodeIndex += degreeDistribution[i]
		i++
	}

	if nodeIndex-numNodes/2 == 1 {
		//
	} else {
		//
	}
	return 0
}
