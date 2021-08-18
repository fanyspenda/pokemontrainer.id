package trainers

import (
	"errors"
)

var (
	// ErrInvalidInput for invalid parameter
	ErrInvalidInput = errors.New("invalid input")

	// ErrGetTrainerFailed error fetching data
	ErrGetTrainerFailed = errors.New("get trainer data error")

	// ErrAddLog error for logging
	ErrAddLog = errors.New("error adding login data to log table")

	// ErrGenerateToken error for logging
	ErrGenerateToken = errors.New("Fail to generate token")

	// ErrInvalidID ...
	ErrInvalidID = errors.New("invalid ID")
)
