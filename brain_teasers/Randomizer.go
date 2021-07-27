package brain_teasers

import (
	"math/rand"
	//"time"
)

// Interface describing all methods a randomizer must implement.
type Randomizer interface {
	Next() int32
}

type intRandomizer struct {
	max int32

	rng *rand.Rand
}

func (ir *intRandomizer) Next() int32 {
	return ir.rng.Int31n(ir.max)
}

// Create a new randomizer
func NewRandomizer() Randomizer {
	rtn := new(intRandomizer)

	//n := time.Now().UnixNano()

	// Locked for consistency.  Switch with the commented line above for a
	// different sequence of numbers each time.
	n := int64(42)

	s := rand.NewSource(n)

	rtn.rng = rand.New(s)

	rtn.max = int32(MAX_INT_16)

	return (rtn)
}
