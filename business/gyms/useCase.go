package gyms

import (
	"context"
	"time"
)

// GymUseCase type of NeewGymUseCase
type GymUseCase struct {
	Repositories   Repositories
	ContextTimeout time.Duration
}

// NewGymUseCases collection of Use Case connected with repositories and timeout
func NewGymUseCases(newRepository Repositories, timeout time.Duration) UseCases {
	return &GymUseCase{
		Repositories:   newRepository,
		ContextTimeout: timeout,
	}
}

// AddGym Case for add new Gym
func (useCase *GymUseCase) AddGym(ctx context.Context, name, address string) (Domain, error) {
	gym, err := useCase.Repositories.AddGym(ctx, name, address)

	if err != nil {
		return Domain{}, err
	}
	return gym, nil
}
