package set

/*
// CoordinateSet is a set implementation geared towards tracking cartesian coordinates
type CoordinateSet struct {
	data         []uint32
	xMax         uint32
	yMax         uint32
	elementCount uint32
	xIter        uint32
	yIter        uint32
}

// NewCoordinateSet creates an empty CoordinateSet with the given boundaries
func NewCoordinateSet(xMax, yMax uint32) CoordinateSet {
	lenData := (xMax * yMax) / 32
	if (xMax*yMax)%32 != 0 {
		lenData++
	}
	return CoordinateSet{xMax: xMax, yMax: yMax,
		data: make([]uint32, lenData), elementCount: 0}
}

// Add a coordinate to the set
func (s *CoordinateSet) Add(x, y uint32) {
	integer, bit := (*s).coordToIntBit(x, y)
	newEntry := uint32(1 << bit)
	currentEntries := (*s).data[integer]
	if (newEntry^math.MaxUint32)&currentEntries == currentEntries {
		(*s).data[integer] += newEntry
		(*s).elementCount++
	}
}

// Contains returns true if the set contains the given coordinate
func (s *CoordinateSet) Contains(x, y uint32) bool {
	integer, bit := (*s).coordToIntBit(x, y)
	query := uint32(1 << bit)
	currentEntries := (*s).data[integer]
	return (query^math.MaxUint32)&currentEntries == currentEntries
}

// Next is used for looping through a set. It returns a coordinate pair and true if the coordinates are valid.
// It returns uint32Max, uint32Max, false if the end of the set was reached without finding an element.
func (s *CoordinateSet) Next() (uint32, uint32, bool) {
	for i := (*s).xIter; i < (*s).xMax; i++ {
		for j := (*s).yIter + 1; j < (*s).yMax; j++ {
			if (*s).Contains(i, j) {
				(*s).xIter = i
				(*s).yIter = j
				return i, j, true
			}
		}
	}
	return uint32(math.MaxUint32), uint32(math.MaxUint32), false
}

// ResetIterator resets the internal iterators to the beginning of the set
func (s *CoordinateSet) ResetIterator() {
	(*s).xIter = 0
	(*s).yIter = 0
}

//Remove an entry
func (s *CoordinateSet) Remove(x, y uint32) {
	integer, bit := (*s).coordToIntBit(x, y)
	toRemove := uint32(1 << bit)
	(*s).data[integer] -= toRemove
}

// Convert a coordinate to a location inside the data array
func (s *CoordinateSet) coordToIntBit(x, y uint32) (int, uint32) {
	integer := ((*s).xMax*x + y) / 32
	bit := ((*s).xMax*x + y) % 32
	return int(integer), bit
}
*/
