package sample

import (
	"errors"
	r "math/rand"
)

// RandGenerator is a generator of random values.
type RandGenerator interface {
	RandInt(min, max int) (int, error)
}

// NewRandGenerator returns a new rand generator.
func NewRandGenerator() RandGenerator {
	return &randGenerator{}
}

type randGenerator struct {
}

func (g *randGenerator) RandInt(min, max int) (int, error) {
	if min > max {
		return 0, errors.New("min should be smaller than max")
	}

	return r.Intn(max-min) + min, nil
}
