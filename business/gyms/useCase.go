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
	if name == "" || address == "" {
		return Domain{}, ErrInvalidInput
	}
	gym, err := useCase.Repositories.AddGym(ctx, name, address)

	if err != nil {
		return Domain{}, ErrAddNewGym
	}
	return gym, nil
}

// UpdateGym update Gym UseCase to run Repository
func (useCase *GymUseCase) UpdateGym(ctx context.Context, gymID int, name, address string) (Domain, error) {
	if name == "" || address == "" {
		return Domain{}, ErrInvalidInput
	}
	if gymID <= 0 {
		return Domain{}, ErrInvalidID
	}

	gym, err := useCase.Repositories.UpdateGym(ctx, gymID, name, address)

	if err != nil {
		return Domain{}, err
	}

	return gym, nil
}

// GetGyms show gyms list. usecase to run repository
func (useCase *GymUseCase) GetGyms(ctx context.Context) ([]Domain, error) {
	gyms, err := useCase.Repositories.GetGyms(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return gyms, nil
}
