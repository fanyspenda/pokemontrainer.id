package trainers

import (
	"context"
	"time"
)

// struct dibawah berfungsi untuk membuat

// TrainerUseCase ...
type TrainerUseCase struct {
	Repository     Repository
	ContextTimeOut time.Duration
}

// NewTrainerUseCase ...
func NewTrainerUseCase(newRepository Repository, timeout time.Duration) UseCase {
	return &TrainerUseCase{
		Repository:     newRepository,
		ContextTimeOut: timeout,
	}
}

// Register as Trainer
func (useCase *TrainerUseCase) Register(ctx context.Context, name, address, username, password string) (Domain, error) {
	result, err := useCase.Repository.Register(ctx, name, address, username, password)

	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

// GetTrainers get all trainers
func (useCase *TrainerUseCase) GetTrainers(ctx context.Context) ([]Domain, error) {
	result, err := useCase.Repository.GetTrainers(ctx)

	if err != nil {
		return []Domain{}, err
	}
	return result, err
}

// AddGym register trainer to gym
// func (useCase *TrainerUseCase) AddGym(ctx context.Context, ID, gymID int) (Domain, error) {
// 	result, err := useCase.Repository.AddGym(ctx, ID, gymID)
// 	if err != nil {
// 		return Domain{}, err
// 	}
// 	return result, nil
// }

// Login as Trainer
func (useCase *TrainerUseCase) Login(ctx context.Context, username, password string) (Domain, error) {
	result, err := useCase.Repository.Login(ctx, username, password)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

// CatchPokemon catch pokemon as Trainer
// func (useCase *TrainerUseCase) CatchPokemon(ctx context.Context, ID, pokemonID int) (Domain, error) {
// 	result, err := useCase.Repository.CatchPokemon(ctx, ID, pokemonID)
// 	if err != nil {
// 		return Domain{}, err
// 	}
// 	return result, nil
// }
