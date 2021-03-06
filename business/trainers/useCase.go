package trainers

import (
	"context"
	"pokemontrainer/business/pokeballs"
	"pokemontrainer/helpers/middlewares"
	"time"
)

// struct dibawah berfungsi untuk membuat

// TrainerUseCase ...
type TrainerUseCase struct {
	Repository     Repository
	ContextTimeOut time.Duration
	LoginLogRepo   MongodbRepository
}

// NewTrainerUseCase ...
func NewTrainerUseCase(newRepository Repository, timeout time.Duration, loginRepo MongodbRepository) UseCase {
	return &TrainerUseCase{
		Repository:     newRepository,
		ContextTimeOut: timeout,
		LoginLogRepo:   loginRepo,
	}
}

// Register as Trainer
func (useCase *TrainerUseCase) Register(ctx context.Context, name, address, username, password string) (Domain, error) {
	if name == "" ||
		address == "" ||
		username == "" ||
		password == "" {
		return Domain{}, ErrInvalidInput
	}
	result, err := useCase.Repository.Register(ctx, name, address, username, password)

	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

// GetFirstBall add trainer data to join table trainer-pokeball
func (useCase *TrainerUseCase) GetFirstBall(ctx context.Context, trainerID uint) (pokeballs.Domain, error) {
	if trainerID <= 0 {
		return pokeballs.Domain{}, ErrInvalidInput
	}
	result, err := useCase.Repository.GetFirstBall(ctx, trainerID)
	if err != nil {
		return pokeballs.Domain{}, err
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
func (useCase *TrainerUseCase) AddGym(ctx context.Context, trainerID, gymID int) (Domain, error) {
	if trainerID <= 0 || gymID <= 0 {
		return Domain{}, ErrInvalidInput
	}

	result, err := useCase.Repository.AddGym(ctx, trainerID, gymID)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

// Login as Trainer
func (useCase *TrainerUseCase) Login(ctx context.Context, username, password string) (Domain, error) {
	if username == "" || password == "" {
		return Domain{}, ErrInvalidInput
	}
	result, err := useCase.Repository.Login(ctx, username, password)
	if err != nil {
		return Domain{}, err
	}

	result, err = useCase.LoginLogRepo.LoginLog(ctx, result.ID)
	if err != nil {
		return Domain{}, ErrAddLog
	}

	result.Token, err = middlewares.GenerateTokenJWT(result.ID)

	if err != nil {
		return Domain{}, ErrGenerateToken
	}

	return result, nil
}

// CatchPokemon catch pokemon as Trainer
func (useCase *TrainerUseCase) CatchPokemon(ctx context.Context, ID, pokemonID int) (Domain, error) {
	if pokemonID <= 0 {
		return Domain{}, ErrInvalidInput
	}
	result, err := useCase.Repository.CatchPokemon(ctx, ID, pokemonID)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

// UpdateTrainer update trainer's profile
func (useCase *TrainerUseCase) UpdateTrainer(ctx context.Context, trainerID int, name, address, username, password string) (Domain, error) {
	if trainerID <= 0 {
		return Domain{}, ErrInvalidID
	}

	if name == "" || address == "" || username == "" || password == "" {
		return Domain{}, ErrInvalidInput
	}
	result, err := useCase.Repository.UpdateTrainer(ctx, trainerID, name, address, username, password)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
