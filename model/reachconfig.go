package model

// ReachType is whatever data type that is used to store agent reaches on the grid
type ReachType = uint8

// ReachConfig is used by makeAgentReaches
type ReachConfig struct {
	// Quantity of the agents
	Quantity uint32
	// Reach of the agents
	Reach ReachType
}
