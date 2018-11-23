package abstractions

import (
	"errors"
	"fmt"
)

const (
	outOfRangeFStr string = "position %v out of range for vector of size %v"
)

// VectorClock encapsulates values
type VectorClock struct {
	vector []int
}

// NewVectorClock creates a vector clock with the given size
func NewVectorClock(size int) VectorClock {
	return VectorClock{make([]int, size)}
}

// SetNewVectorClock creates a vector clocks with arbitrary inital values
func SetNewVectorClock(vals []int) VectorClock {
	return VectorClock{vals}
}

// Size gives number of values in vector
func (c *VectorClock) Size() int {
	return len(c.vector)
}

// Set allows explicit setting of a value in the vector
func (c *VectorClock) Set(pos, newVal int) error {
	if !c.inRange(pos) {
		return fmt.Errorf(outOfRangeFStr, pos, c.Size())

	} else if curVal, _ := c.At(pos); newVal < curVal {
		return fmt.Errorf("setting vector to %v not allowed because it is lesser than currrent value %v", newVal, curVal)

	}
	c.vector[pos] = newVal
	return nil
}

// At returns value of vector at spec'd position
func (c *VectorClock) At(pos int) (int, error) {
	if !c.inRange(pos) {
		return -1, fmt.Errorf(outOfRangeFStr, pos, c.Size())
	}
	return c.vector[pos], nil
}

// Incr increments value at specified pos
func (c *VectorClock) Incr(pos int) error {
	if !c.inRange(pos) {
		return errors.New("cannot increment vector at pos=%v")
	}
	c.vector[pos]++
	return nil
}

// Equals return true iff both vectors are same size and contain same values in each position
func (c *VectorClock) Equals(otherClock VectorClock) bool {
	if c.Size() != otherClock.Size() {
		return false
	}

	for pos, val := range c.vector {
		if otherVal, _ := otherClock.At(pos); val != otherVal {
			return false
		}
	}
	return true
}

func (c *VectorClock) inRange(pos int) bool {
	if pos < 0 || pos >= c.Size() {
		return false
	}
	return true
}

func (c VectorClock) String() string {
	return fmt.Sprintf("%v", c.vector)
}
