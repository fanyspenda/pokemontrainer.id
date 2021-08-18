package gyms

import (
	"errors"
)

var (
	// ErrInvalidInput for invalid input
	ErrInvalidInput = errors.New("invalid input")

	// ErrAddNewGym for input correct, but cannot add new Gym
	ErrAddNewGym = errors.New("Failed to add new Gym")

	// ErrInvalidID for id under 0
	ErrInvalidID = errors.New("invalid ID")
)
