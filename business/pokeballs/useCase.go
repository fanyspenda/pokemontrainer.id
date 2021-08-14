package pokeballs

import (
	"context"
	"time"
)

// PokeballUseCase to make newPokeballUseCase
type PokeballUseCase struct {
	Repositories   Repositories
	TimeoutContext time.Duration
}

// NewPokeballUseCase to be used in controller
func NewPokeballUseCase(newRepositories Repositories, newTimeout time.Duration) UseCases {
	return &PokeballUseCase{
		Repositories:   newRepositories,
		TimeoutContext: newTimeout,
	}
}

// AddPokeball add pokeball usecase
func (useCase *PokeballUseCase) AddPokeball(ctx context.Context, name string, successRate float32) (Domain, error) {
	domain, err := useCase.Repositories.AddPokeball(ctx, name, successRate)
	if err != nil {
		return Domain{}, err
	}
	return domain, nil
}
